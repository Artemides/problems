const headerButton = document.getElementById("sort-title");

const API_URL = "http://localhost:3000";

const sortBy = (param) => {
  const url = `${window.location.pathname}?sortBy=${param}`;
  window.history.pushState({}, "", url);
  window.location.reload();
};

const sortButtons = document.querySelectorAll(".header-btn");
sortButtons.forEach((sortButton) => {
  sortButton.addEventListener("click", (event) => {
    const by = event.target.getAttribute("data-sortBy");
    sortBy(by);
  });
});

window.addEventListener("load", () => {
  const url = new URL(window.location.href);
  const by = url.searchParams.get("sortBy");
  if (!by) {
    return;
  }

  const selectedButton = document.getElementById(`sort-${by.toLowerCase()}`);
  selectedButton.style.color = "#f04722";
  selectedButton.style.borderBottom = "1px solid #f04722";
});
