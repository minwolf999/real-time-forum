import { logIn } from "./login.js";
import { Home } from "./home/home.js";
import { HomeCenter } from "./home/homecenter.js";
import { Profile } from "./profile/Profile.js";
import { Post } from "./post/post.js";
import { WriteMessage } from "./chatwithsomeone.js";

export const socket = new WebSocket("ws://localhost:8080/login")
export const pageName = { pagename: 'login' }
export var User

/**
 * Launch the LogIn function at the load of the page
 */
document.addEventListener('DOMContentLoaded', () => {
    logIn()
})

/**
 * At the unloading of the page
 * Close the websocket at if the user was connected disconnect him
 */
window.onbeforeunload = () => {
    socket.close()

    if (User.hasOwnProperty('Username')) {
        var xhr = new XMLHttpRequest();
        xhr.open('POST', '/disconnect05842365165', false);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]));
    }
}

/**
 * manage the receives of datas in the websocket
 * 
 * @param {Any} e contain the json of the data
 */
socket.onmessage = (e) => {
    var updatedData = JSON.parse(e.data);
    if (updatedData.hasOwnProperty('login')) {
        if (updatedData.login.hasOwnProperty('error')) {
            const verif = document.getElementById('error')
            if (verif != null) {
                verif.remove()
            }

            const p = document.createElement('p')
            p.id = "error"
            p.innerHTML = updatedData.login.error
            
            document.getElementsByTagName('center')[0].append(p)
        } else if (updatedData.login.hasOwnProperty('success')) {
            User = updatedData.login.success

            Home()
        }
    }

    if (updatedData.hasOwnProperty("register")) {
        if (updatedData.register.hasOwnProperty("error")) {
            const verif = document.getElementById('error')
            if (verif != null) {
                verif.remove()
            }

            const p = document.createElement('p')
            p.id = "error"
            p.innerHTML = updatedData.register.error

            document.getElementsByTagName('center')[0].append(p)
        } else if (updatedData.register.hasOwnProperty('success')) {
            User = updatedData.register.success

            Home()
        }
    }

    if (updatedData.hasOwnProperty("home")) {
        if (updatedData.home.hasOwnProperty("error")) {
            const p = document.createElement('p')
            p.innerHTML = updatedData.home.error

            document.getElementById('error').className = "error"
            document.getElementById('error').innerHTML = p
        } else if (updatedData.home.hasOwnProperty('success')) {
            HomeCenter()
        }
    }

    if (updatedData.hasOwnProperty("filter")) {
        if (updatedData.filter.hasOwnProperty("error")) {
            const p = document.createElement('p')
            p.innerHTML = updatedData.filter.error

            document.getElementById('error').className = "error"
            document.getElementById('error').innerHTML = p
        } else if (updatedData.filter.hasOwnProperty('success')) {
            Home(true)
        }
    }

    if (updatedData.hasOwnProperty("profile")) {
        Profile(true)
    }

    if (updatedData.hasOwnProperty('modification')) {
        if (updatedData.modification.hasOwnProperty('error')) {
            const verif = document.getElementById('error')
            if (verif != null) {
                verif.remove()
            }

            const p = document.createElement('p')
            p.id = 'error'
            p.innerHTML = updatedData.modification.error
            
            document.getElementsByTagName('center')[0].append(p)
        } else if (updatedData.modification.hasOwnProperty('success')) {
            User = updatedData.modification.success
            Profile()
        }
    }

    if (updatedData.hasOwnProperty('comment')) {
        if (updatedData.comment.hasOwnProperty('error')) {
            const p = document.createElement('p')
            p.innerHTML = updatedData.filter.error

            document.getElementById('error').className = "error"
            document.getElementById('error').innerHTML = p
        } else if (updatedData.comment.hasOwnProperty('success')) {
            let datas = updatedData.comment.success.split('|')

            if (datas[0] == "success") {
                Post(datas[1], 0)
            } else {
                Post(datas[1], 1)
            }
        }
    }

    if (updatedData.hasOwnProperty('chat') && pageName.pagename === "chatWS") {
        if (updatedData.chat.hasOwnProperty('error')) {
            const p = document.createElement('p')
            p.innerHTML = updatedData.chat.error

            document.getElementById('error').className = "error"
            document.getElementById('error').innerHTML = p
        } else if (updatedData.chat.hasOwnProperty('success') && pageName.interlocutor === updatedData.chat.success[0] && updatedData.chat.success.length == 1) {
            document.getElementsByClassName('previousMessage')[0].innerHTML = ''
            WriteMessage(updatedData.chat.success[0])
            document.getElementsByClassName('previousMessage')[0].scrollTop = document.getElementsByClassName('previousMessage')[0].scrollHeight;
        } else if (updatedData.chat.success.length > 1) {
            document.getElementById('isWritting').style.display = "inline"
            const pointDiv = document.createElement('div')
            pointDiv.id = 'pointDiv'
            pointDiv.className = 'pointDiv'
            document.getElementById('isWritting').appendChild(pointDiv)

            const timer = setInterval(() => {
                const elem = document.getElementById('pointDiv')

                if (elem.innerHTML == "") {
                    elem.innerHTML = "."
                } else if (elem.innerHTML == ".") {
                    elem.innerHTML = ".."
                } else if (elem.innerHTML == "..") {
                    elem.innerHTML = "..."
                } else if (elem.innerHTML == "...") {
                    elem.innerHTML = ""
                }
            }, 100)

            setTimeout(() => {
                clearInterval(timer)
                document.getElementById('pointDiv').remove()
                document.getElementById('isWritting').style.display = "none"
            }, 1000)
        }
    }

    if (updatedData.hasOwnProperty('chat') && (pageName.pagename === "profile" || pageName.pagename === "home" ||pageName.pagename === "chatWS" && pageName.interlocutor !== parseInt(updatedData.chat.success))) {
        document.getElementById('notification').textContent++
    }
}
