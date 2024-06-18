import { Home } from "./home/home.js";
import { User, pageName, socket } from "./index.js"
import { Profile } from "./profile/Profile.js";
import { SetNewCSS } from "./setnewcss.js";
import { throttle } from "./throttle.js";

let maxMessage = 10;

/**
 * manage the chat with someone page
 */
export const ChatWithSomeone = async (interlocutorId, interlocutorName = undefined) => {
    pageName.pagename = 'chatWS'
    pageName.interlocutor = interlocutorId
    SetNewCSS(['/template/css/chatwithsomeone.css'])
    
    const notification = await fetch('http://localhost:8080/getnotificationQuantity05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    document.getElementById('content').innerHTML = `
    <div class="barre_menu">
        <button type="button" id='home' class="home">Home</button>

        <h2 class="text">Chat with ${interlocutorName ? interlocutorName : data[0].Sender.Username !== User.Username ? data[0].Sender.Username : data[0].Recever.Username}</h2>

        <div class="profileDiv">
            <button type="button" id="profile" class="profile">Profile</button>
            <div id="notification" class="notification">${notification}</div>
        </div>
    </div>

    <div class="body">
        <div class="previousMessage"></div>

        <div class="form">
            <div id="isWritting" class="isWritting" style="display: none">${interlocutorName} is writting</div>
            <textarea id="textarea" placeholder="Write your message"></textarea>
            <button id="valider">Envoyer le message</button>
            
            <div id="error"></div>
        </div>
    </div>
    `
    document.getElementsByClassName('previousMessage')[0].innerHTML = ''
    WriteMessage(interlocutorId)
    document.getElementsByClassName('previousMessage')[0].scrollTop = document.getElementsByClassName('previousMessage')[0].scrollHeight;

    document.getElementById('home').addEventListener('click', () => {
        Home()
    })

    document.getElementById('profile').addEventListener('click', () => {
        Profile()
    })
    
    document.getElementById('textarea')
        .addEventListener('keydown', throttle(() => {
            let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'isWritting',
                senderId: User.Id,
                interlocutorId: interlocutorId,
            }

            socket.send(JSON.stringify(formResponse));
        }, 1000))

    document.getElementById('valider').addEventListener('click', () => {
        let formResponse = {}

        formResponse[pageName.pagename] = {
            senderId: User.Id,
            interlocutorId: interlocutorId,

            message: document.getElementById('textarea').value,
        }

        socket.send(JSON.stringify(formResponse));

        document.getElementById('textarea').value = ""
    })

    document.getElementsByClassName('previousMessage')[0].addEventListener('scroll', () => {
        const elem = document.getElementsByClassName('previousMessage')[0]

        if (elem.scrollHeight <= Math.abs(elem.scrollTop) + elem.clientHeight) {
            maxMessage += 10
            WriteMessage(interlocutorId)
        }
    })
}

/**
 * manage the message zone
 * 
 * @param {string} interlocutorId 
 */
export const WriteMessage = async (interlocutorId) => {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/clearNotification05842365165', false);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password, interlocutorId]));

    const data = await fetch('http://localhost:8080/getmessages05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password, interlocutorId, maxMessage.toString()]) }).then(response => response.json())

    if (data) {
        data.forEach(message => {
            const div = document.createElement('div')
            div.className = message.Sender.Username === User.Username ? "sender" : "recever"
    
            div.innerHTML = `
                <div class="firstLine">
                    <p>${message.Sender.Username} - ${message.Date}</p>
                </div>
                <p>${message.Message}</p>
            `
    
            document.getElementsByClassName('previousMessage')[0].appendChild(div)
        });
    }
}