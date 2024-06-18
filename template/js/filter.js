import { Home } from "./home/home.js";
import { socket, pageName, User } from "./index.js";
import { SetNewCSS } from "./setnewcss.js";

/**
 * manage the filter page
 */
export const Filter = () => {
    pageName.pagename = 'filter'
    SetNewCSS(['/template/css/filtre.css']);
    
    document.getElementById('content').innerHTML = `
        <div class="barre_menu">
            <button type="button" class="home" id="filterButton">Home</button>

            <h2 class="text">Filter Page</h2>

            <p></p>
        </div>

        <div class="body" id="body">
            <div class="bodyLeft"></div>
            <div class="bodyCenter"></div>
            <div class="bodyRight"></div>
        </div>

        <img src="/template/image/div_cote.png" class="img1">
        <img src="/template/image/div_cote.png" class="img2">
        <img src="/template/image/eyes.png" class="img3">

        <div class="filterName">
            <span>Filter by namePost</span><br>
            <input type="text" id="namePost" placeholder="Name of a post"><br>
            <input type="submit" id="filterByName" value="Valider"><br>
        </div>

        <div id="error"></div>

        <center>
            <div class="likes">
                <button id="likes" class="filtreLike">Filter by Likes</button>
                <button id="dislikes" class="filtreDislike">Filter by Dislikes</button>
            </div>
        </center>
    `

    document.getElementById('filterButton').addEventListener('click', () => {
        Home()
    })

    document.getElementById('filterByName').addEventListener('click', () => {
        let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filterByName',
                value: document.getElementById('namePost').value,
            }

            socket.send(JSON.stringify(formResponse));
    })

    document.getElementById('likes').addEventListener('click', () => {
        let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filterByLikes',
                value: 'filterByLikes',
            }

            socket.send(JSON.stringify(formResponse));
    })

    document.getElementById('dislikes').addEventListener('click', () => {
        let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filterByDislikes',
                value: 'filterByDislikes',
            }

            socket.send(JSON.stringify(formResponse));
    })
}
