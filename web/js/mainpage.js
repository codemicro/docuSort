String.prototype.insertAt = function(i, str) {
    return this.slice(0, i) + str + this.slice(i);
}

let allTopics = []
let newList = document.createElement("ol")

for (let i = 0; i < items.length; i++) {
    let currentItem = items[i]
    
    var a = document.createElement("a")
        
    var filename = currentItem["Filename"].split("\\")[1] // splitting it because it contains the path
    var splitFilename = filename.split(" ")

    var date = splitFilename[1].insertAt(2, "/").insertAt(5, "/")

    a.dataset.document = filename
    a.dataset.topics = JSON.stringify(currentItem["Topics"])
    
    a.addEventListener("mouseup", showModal);


    for (let x = 0; x < currentItem["Topics"].length; x++) {
        if (!allTopics.includes(currentItem["Topics"][x])) {
            allTopics.push(currentItem["Topics"][x])
        }
    }

    a.setAttribute("href", "#")
    a.appendChild(document.createTextNode(date + " - " + currentItem["Topics"].join(", ")))

    var li = document.createElement("li")
    li.appendChild(a)

    newList.appendChild(li)

}

allTopics = allTopics.sort()
document.getElementById("buildArea").appendChild(newList)