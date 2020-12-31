window.addEventListener(
    "onload",
    (e) => {

    }
);

const toggleDarkMode = ()=>{
    if (document.querySelector("html").dataset["theme"] == "light") {
        document.querySelector("html").dataset["theme"] = "dark";
    } else {
        document.querySelector("html").dataset["theme"] = "light";
    }
};
