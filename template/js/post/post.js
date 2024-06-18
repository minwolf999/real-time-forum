import { Home } from "../home/home.js";
import { User, pageName, socket } from "../index.js";
import { Profile } from "../profile/Profile.js";
import { SetNewCSS } from "../setnewcss.js";
import { Comments } from "./comments.js";
import { messageSender } from "./messagesender.js";
import { postInformation } from "./postInformation.js";

/**
 * manage the post page
 * @param {string} id 
 * @param {number} filtred 
 */
export const Post = async (id, filtred = 0) => {
    pageName.pagename = 'comment'
    SetNewCSS(['/template/css/post/post.css', '/template/css/post/comments.css',
    '/template/css/post/messageSender.css', '/template/css/post/postInformation.css'
    ])

    const data = await fetch('http://localhost:8080/getcomments05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password, id, filtred.toString()]) }).then(response => response.json())

    document.getElementById('content').innerHTML = `
        <div class="barre_menu">
            <button type="button" class="home" id="home">Main Page</button>

            <h2 class="text">${data.Posts[0].Name}</h2>

            <button type="button" class="profile" id="profile">
                    ${User.Username}
            </button>
        </div>

        <div class="body" id="body">
            <div class="bodyLeft" id="bodyLeft"></div>
            <div class="bodyCenter" id="bodyCenter"></div>
            <div>
                <div class="bodyRight" id="bodyRight"></div>

                <div class="filter">
                    <p>Filter by</p>
                    <div>
                        <button type="submit" name="filtre" value="like">Like</button>
                        <button type="submit" name="filtre" value="dislike">Dislike</button>
                    </div>
                </div>
            </div>
        </div>
    `

    document.getElementById("home").addEventListener('click', () => {
        Home()
    })

    document.getElementById('profile').addEventListener('click', () => {
        Profile()
    })

    document.getElementsByName('filtre').forEach(button => {
        button.addEventListener('click', () => {
            let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filtre',
                value: button.value,
                id: data.Posts[0].Id,
            }
    
            socket.send(JSON.stringify(formResponse));
        })
    })

    postInformation(data.Posts[0])
    Comments(data.Comments, data.Posts[0].Id)
    messageSender(data.Posts[0].Id)
}