import { socket, pageName, User } from "../index.js";

/**
 * manage the display of the filter by categorie
 */
export const HomeLeft = async () => {
    let data = await fetch('http://localhost:8080/getcategories05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    document.getElementsByClassName('bodyLeft')[0].innerHTML = `
        <div id="form">
            <button name="categories" value="All" class="categorie">All Categories</button>
        </div>
        <br>
    `

    data['Categories'].forEach(element => {
        let button = document.createElement('button')
        button.setAttribute('name', 'categories')
        button.setAttribute('value', element.Id)
        button.setAttribute('class', 'categorie')
        
        button.textContent = element.Name

        document.getElementById('form').append(button)
    });

    let buttons = Array.from(document.getElementsByName('categories'))
    buttons.forEach(button => {
        button.addEventListener('click', () => {
            let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filter',
                filterType: button.value
            }

            socket.send(JSON.stringify(formResponse));
        })
    })
}
