String.prototype.insertAt = function(i, str) {
    return this.slice(0, i) + str + this.slice(i);
}

Array.prototype.sortOn = function(key){
    this.sort(function(a, b){
        if(a[key] < b[key]){
            return -1;
        }else if(a[key] > b[key]){
            return 1;
        }
        return 0;
    });
}

items.sortOn("Filename")

let allTopics = []
// Create list of items
function createListOfItems(filter="") {
    const newList = document.createElement("ol")
    for (let i = 0; i < items.length; i++) {
        let currentItem = items[i]
        
        var a = document.createElement("a")
            
        var filename = currentItem["Filename"].split("\\")[1] // splitting it because it contains the path
        var splitFilename = filename.split(" ")
    
        var date = splitFilename[0]
    
        a.dataset.document = filename
        a.dataset.topics = JSON.stringify(currentItem["Topics"])
        
        a.addEventListener("mouseup", showModal);
    
        a.setAttribute("href", "#")
        a.appendChild(document.createTextNode(date + " - " + currentItem["Teacher"] + ", " + currentItem["Type"] + " - " + currentItem["Topics"].join(", ")))
    
        var li = document.createElement("li")
        li.appendChild(a)
    
        if (filter == "" || currentItem["Topics"].includes(filter)) {
            newList.appendChild(li)
        }
    }
    return newList
}

for (let i = 0; i < items.length; i++) {
    let currentItem = items[i]
    for (let x = 0; x < currentItem["Topics"].length; x++) {
        if (!allTopics.includes(currentItem["Topics"][x])) {
            allTopics.push(currentItem["Topics"][x])
        }
    }
}

// Populate list of dropdowns
allTopics = allTopics.sort()
const dropdownMenuItems = document.getElementById("dropdownMenuItems")
const dropdownMenuStatus = document.getElementById("dropdownMenuDescription")
const buildArea = document.getElementById("buildArea")

function filter(sender) {
    selectedTopic = sender.target.dataset.topic
    dropdownMenuStatus.innerText = selectedTopic

    const filteredList = createListOfItems(selectedTopic)

    buildArea.replaceChild(filteredList, buildArea.children[0])

}

for (let i = 0; i < allTopics.length; i++) {
    // <a class="dropdown-item" href="#">Something else here</a>
    var a = document.createElement("a")

    a.dataset.topic = allTopics[i]
    a.href = "#"
    a.innerText = allTopics[i]
    a.classList.add("dropdown-item")

    a.addEventListener("mouseup", filter)

    dropdownMenuItems.appendChild(a)
}

buildArea.appendChild(createListOfItems())