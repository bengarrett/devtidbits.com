"use strict";

function blueText() {
    var txt = document.getElementById("text");
    txt.style = "color: blue"; // throws an error in IE and Edge
    //txt.style.cssText = "color: blue"; // works correctly in all browsers
}
