import { socket, pageName, User } from "../index.js";
import { SetNewCSS } from "../setnewcss.js";
import { Profile } from "./Profile.js";

/**
 * manage the modification profile page
 */
export const Modification = () => {
    pageName.pagename = 'modification'
    SetNewCSS(['/template/css/modificationUser.css'])

    document.getElementById('content').innerHTML = `
        <center>
            <h2>Modification</h2>
            <button class="profile" id="profile">Profile</button>

            <img src="/template/image/div_cote.png" class="img1">
            <img src="/template/image/div_cote.png" class="img2">
            <img src="/template/image/eyes.png" class="img3">
            
            <div id="form" class="form">
                <div class="username">
                    <span>Username</span><br>
                    <input type="text" id="username" maxlength="20" value="${User.Username}">
                </div>

                <div class="age">
                    <span>Age</span><br>
                    <input type="number" id="age" value="${User.Age}">
                </div>

                <div class="gender">
                    <span>Gender</span><br>
                    <select id="gender">
                        <option value="Male">Male</option>
                        <option value="Female">Female</option>
                    </select>
                </div>

                <div class="firstName">
                    <span>firstName</span><br>
                    <input type="text" id="firstName" maxlength="20" value="${User.FirstName}">
                </div>

                <div class="LastName">
                    <span>lastName</span><br>
                    <input type="text" id="lastName" maxlength="20" value="${User.LastName}">
                </div>

                <div class="email">
                    <span>Email</span><br>
                    <input type="email" id="email" value="${User.Email}">
                </div>

                <div class="password">
                    <span>Password</span><br>
                    <input type="password" id="password">
                </div>

                <input type="submit" id="profile" value="Profile Page">
                <button id="valide" class="login">Validate</button>
            </div>
            
        </center>
    `

    document.getElementById('profile').addEventListener('click', () => {
        Profile()
    })

    document.getElementById("valide").addEventListener("click", () => {
        let formResponse = {}
            
        formResponse[pageName.pagename] = {
            type: 'modification',
            currentUsername: User.Username,
            currentAge: User.Age,
            currentGender: User.Gender,
            currentFirstName: User.FirstName,
            currentLastName: User.LastName,
            currentEmail: User.Email,
            currentPassword: User.Password,

            id: User.Id,
            username: document.getElementById('username').value,
            age: document.getElementById('age').value.toString(),
            gender: document.getElementById('gender').value,
            firstName: document.getElementById('firstName').value,
            lastName: document.getElementById('lastName').value,
            email: document.getElementById('email').value,
            password: document.getElementById('password').value,
        }

        socket.send(JSON.stringify(formResponse));
    })
}
