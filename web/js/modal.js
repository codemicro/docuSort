function showModal(event) {
    const src = event.target.dataset.document;
    console.log(src)
  
    const newImage = document.createElement("iframe")
    newImage.src = src;
    newImage.id = "modalIframe";
    
    modal.querySelector("#modalIframe").replaceWith(newImage);
  
    modal.style.display = "block";
}

const modal = document.getElementById("displayModal");
const modalCloseButton = document.getElementById("closeButton");

// Close modal if item with the ID photoModal is clicked
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
}

modalCloseButton.onclick = function() {
  modal.style.display = "none";
}

document.addEventListener("keyup", function (event) {
  if(event.key == "Escape") {
    modal.style.display = "none";
  }
});
