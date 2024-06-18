import { SetNewCSS } from "./setnewcss.js";
import { register } from "./register.js";

import { socket, pageName } from "./index.js";

/**
 * manage the logIn page
 */
export const logIn = () => {
    pageName.pagename = 'login'
    SetNewCSS(['/template/css/login.css'])

    document.getElementById('content').innerHTML = `
        <center id="centerTag">
            <h2>Log In</h2>
            
            <img src="/template/image/div_cote.png" class="img1">
            <img src="/template/image/div_cote.png" class="img2">
            <img src="/template/image/eyes.png" class="img3">
            
            <div class="form">
                <div class="email">
                    <span>Email</span><br>
                    <input type="text" id="email">
                </div>

                <div class="password">
                    <span>Password</span><br>
                    <input type="password" id="password">
                </div>

                <input type="submit" id="login_confirmation" value="Validate">
                <button type="button" id="register_redirect" class="register">Register</button>
            </div>
        </center>
    `

    document.getElementById('login_confirmation').addEventListener('click', () => {
        let formResponse = {}

        formResponse[pageName.pagename] = {
            email: document.getElementById('email').value,
            password: document.getElementById('password').value,
        }

        socket.send(JSON.stringify(formResponse));
    })

    document.getElementById('register_redirect').addEventListener('click', () => {
        register()
    })
}