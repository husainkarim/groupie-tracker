const image = document.querySelector("img");

image.setAttribute("cache-control", "max-age=2592000");
image.setAttribute("expires", "Thu, 01 Jan 2024 00:00:00 GMT");

const value1 = document.getElementById("value1");
const input1 = document.getElementById("year-selector-start");

input1.addEventListener("input", (event) => {
    value1.textContent = input1.value;
    input2.setAttribute("min", input1.value);
});

const value2 = document.getElementById("value2");
const input2 = document.getElementById("year-selector-end");

input2.addEventListener("input", (event) => {
    value2.textContent = input2.value;
});

const start = document.getElementById("year-start");
const end = document.getElementById("year-end");

start.addEventListener("input", (event) => {
    end.setAttribute("min", start.value);
});
