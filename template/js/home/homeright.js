import { socket, pageName, User } from "../index.js";

/**
 * manage the display of the send post zone
 */
export const HomeRight = async () => {
    let data = await fetch('http://localhost:8080/getcategories05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    if (User.Username !== 'visitor') {
        document.getElementsByClassName('bodyRight')[0].innerHTML = `
            <div class="messageZone">
                <div if='form'>
                    <center>
                        <span class="description0">Categories of the new Post</span>
                    </center>
    
                    <select name="sendCategorie" class="selectCategorie" id="select" multiple>
                    </select><br>

                    <center>
                        <span class="description1">Name of the new Post</span>
                    </center>
                    <input type="text" id="sendWriteMessage" class="sendWriteMessage" placeholder="Write the name of your Post">
            
                    <center>
                        <span class="description2">Description of the new Post</span>
                    </center>
                    <textarea id="sendWriteDescription" class="sendWriteDescription" placeholder="Write the description of your Post"></textarea>
            
                    <button id="poster" class="poster">Send</button>

                    <div id="error"></div>
                </div>
            </div>
        `

        data.Categories.forEach(categorie => {
            const option = document.createElement("option")
            option.value = categorie.Id
            option.innerHTML = categorie.Name

            document.getElementById('select').appendChild(option)
        });

        document.getElementById('poster').addEventListener('click', () => {
            let formResponse = {}
            
            let arrayHTMLElement = Array.from(document.getElementById('select').selectedOptions)
            let idCategories = []

            arrayHTMLElement.forEach(elem => {
                idCategories.push(elem.value)
            })

            formResponse[pageName.pagename] = {
                type: 'post',
                idCreator: data.User.Id,
                categorie: idCategories.join('|'),
                message: document.getElementById('sendWriteMessage').value,
                description: document.getElementById('sendWriteDescription').value,
            }

            socket.send(JSON.stringify(formResponse));
        })
    }
}
