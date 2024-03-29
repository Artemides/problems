// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 234.

// The thumbnail package produces thumbnail-size images from
// larger images.  Only JPEG images are currently supported.
package concurrency

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Image returns a thumbnail-size version of src.
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

func ThumbnailImages(filenames []string) {
	for _, filename := range filenames {
		if _, err := ImageFile(filename); err != nil {
			log.Printf("thumbling %s image, error: %s", filename, err)
		}
	}
}

func ThumbnailImagesParallel(filenames []string) {
	ch := make(chan struct{})
	for _, filename := range filenames {
		go func(f string) {
			ImageFile(f)
		}(filename)
	}

	for range filenames {
		<-ch
	}
}
func ThumbnailImagesParallel2(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		err       error
		thumbfile string
	}
	ch := make(chan item, len(filenames))
	for _, filename := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(filename)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func ThumbnailImagesWait(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for filename := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()

			thumbfile, err := ImageFile(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()

		}(filename)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total

}

func ThumbnailMain() {

	filenames := []string{"ada.jpeg", "btc.jpeg", "sol.jpeg", "xrp.jpeg", "tether.jpeg"}
	filesChan := make(chan string)
	var wg sync.WaitGroup

	for _, filename := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			fmt.Println("f", f)
			filesChan <- "./assets/" + f
		}(filename)
	}
	go func() {
		wg.Wait()
		close(filesChan)
	}()

	thumbFilesSize := ThumbnailImagesWait(filesChan)
	fmt.Println("total sizes: ", thumbFilesSize)
}
