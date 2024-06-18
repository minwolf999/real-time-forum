import { SetNewCSS } from "../setnewcss.js";
import { Filter } from "../filter.js";
import { HomeLeft } from "./homeleft.js";
import { HomeCenter } from "./homecenter.js";
import { HomeRight } from "./homeright.js";
import { Profile } from "../profile/Profile.js";

import { pageName, User } from "../index.js";

/**
 * manage the home page
 * 
 * @param {boolean} filtered 
 */
export const Home = async (filtered = false) => {
    pageName.pagename = 'home'

    SetNewCSS(['/template/css/home.css', '/template/css/home/bodyCenter.css',
     '/template/css/home/bodyleft.css', '/template/css/home/bodyRight.css']);

    const data = await fetch('http://localhost:8080/getnotificationQuantity05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    document.getElementById('content').innerHTML = `
        <div class="barre_menu">
            <button type="button" class="filter" id="filterButton">Filtrer</button>

            <h2 class="text">Home Page</h2>

            <div class="profileDiv">
                <button type="button" class="profile" id="profileButton">${User["Username"]}</button>
                <div id="notification" class="notification">${data}</div>
            </div>
        </div>

        <div class="body" id="body">
            <div class="bodyLeft"></div>
            <div class="bodyCenter"></div>
            <div class="bodyRight"></div>
        </div>
    `

    document.getElementById('filterButton').addEventListener('click', () => {
        Filter()
    })

    document.getElementById('profileButton').addEventListener('click', () => {
        Profile()
    })

    HomeLeft()
    HomeCenter(filtered)
    HomeRight()
}
