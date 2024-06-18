import { User, pageName } from "./index.js"
import { Home } from "./home/home.js"
import { Profile } from "./profile/Profile.js"
import { SetNewCSS } from "./setnewcss.js"
import { ChatWithSomeone } from "./chatwithsomeone.js"
import { throttle } from "./throttle.js"

/**
 * manage the chat page
 */
export const Chat = async () => {
    pageName.pagename = 'chat'
    SetNewCSS(['/template/css/chat.css'])

    const data = await fetch('http://localhost:8080/getusers05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    document.getElementById('content').innerHTML = `
    <div class="barre_menu">
        <button type="button" id='home' class="home">Home</button>

        <h2 class="text">Peoples Page</h2>

        <button type="button" id="profile" class="profile">Profile</button>
    </div>

    <center>
        <div class="body"></div>
    </center>
    `
    if (data) {
        data.forEach(user => {
            const div = document.createElement('div')
            div.id = user.Id
            div.className = 'user'
    
            div.innerHTML = `
                <h2>${user.Username}</h2>
                <div class="${user.Connected == 0 ? "disconnected" : "connected"}">${user.Notification === 0 ? "" : user.Notification}</div>
            `
            document.getElementsByClassName('body')[0].appendChild(div)
    
            div.addEventListener('click', () => {
                clearInterval(timer)
                ChatWithSomeone(user.Id, user.Username)
            })
        });
    }

    const timer = setInterval(throttle(async () => {
        const data = await fetch('http://localhost:8080/getusers05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

        Array.from(document.getElementsByClassName('user')).forEach(elem => {
            elem.remove()
        })

        if (data) {
            data.forEach(user => {
                const div = document.createElement('div')
                div.id = user.Id
                div.className = 'user'
        
                div.innerHTML = `
                    <h2>${user.Username}</h2>
                    <div class="${user.Connected == 0 ? "disconnected" : "connected"}">${user.Notification === 0 ? "" : user.Notification}</div>
                `
                document.getElementsByClassName('body')[0].appendChild(div)
        
                div.addEventListener('click', () => {
                    clearInterval(timer)
                    ChatWithSomeone(user.Id, user.Username)
                })
            })
        }
    }, 2000), 2000)

    document.getElementById('home').addEventListener('click', () => {
        clearInterval(timer)
        Home()
    })

    document.getElementById('profile').addEventListener('click', () => {
        clearInterval(timer)
        Profile()
    })
}