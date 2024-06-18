import { SetNewCSS } from "./setnewcss.js";
import { logIn } from "./login.js";
import { socket, pageName } from "./index.js";

/**
 * manage the register page
 */
export const register = () => {
    pageName.pagename = 'register'
    SetNewCSS(['/template/css/register.css'])

    document.getElementById('content').innerHTML = `
        <center id="centerTag">
            <h2>Register</h2>

            <img src="/template/image/div_cote.png" class="img1">
            <img src="/template/image/div_cote.png" class="img2">
            <img src="/template/image/eyes.png" class="img3">
            
            <div class="form">
                <div class="username">
                    <span>Username</span><br>
                    <input type="text" id="username" maxlength="20">
                </div>

                <div class="age">
                    <span>Age</span><br>
                    <input type="number" id="age">
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
                    <input type="text" id="firstName" maxlength="20">
                </div>

                <div class="LastName">
                    <span>lastName</span><br>
                    <input type="text" id="lastName" maxlength="20">
                </div>

                <div class="email">
                    <span>Email</span><br>
                    <input type="email" id="email">
                </div>

                <div class="password">
                    <span>Password</span><br>
                    <input type="password" id="password">
                </div>

                <div class="confirm">
                    <span>Confirm Password</span><br>
                    <input type="password" id="confirm">
                </div>

                <input type="submit" id="register_confirmation" value="Validate">
                <button type="button" id="login_redirect" class="login">Log In</button>
            </div>
        </center>
    `

    document.getElementById('register_confirmation').addEventListener('click', () => {
        let formResponse = {}

        formResponse[pageName.pagename] = {
            username: document.getElementById('username').value,
            age: document.getElementById('age').value.toString(),
            gender: document.getElementById('gender').value,
            firstName: document.getElementById('firstName').value,
            lastName: document.getElementById('lastName').value,
            email: document.getElementById('email').value,
            password: document.getElementById('password').value,
            confirm: document.getElementById('confirm').value,
        }

        socket.send(JSON.stringify(formResponse));
    })

    document.getElementById('login_redirect').addEventListener('click', () => {
        logIn()
    })
}