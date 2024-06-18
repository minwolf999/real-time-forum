import { User, pageName, socket } from "../index.js"


/**
 * manage the send message zone
 * 
 * @param {string} postId 
 */
export const messageSender = (postId) => {
    document.getElementById('bodyRight').innerHTML = `
        <div class="messageZone">
            <div class='form'>
                <center>
                    <div id='error'></div>
                </center>

                <center>
                    <span class="description2">Comment writing zone</span>
                </center>
                <textarea id="commentText" class="sendWriteDescription" placeholder="Write the Comment"></textarea>

                <button id='sendComment' type="submit" class="poster">Send</button>
            </div>
        </div>
    `

    document.getElementById('sendComment').addEventListener('click', () => {
        let formResponse = {}

        formResponse[pageName.pagename] = {
            type: 'sendComment',
            text: document.getElementById('commentText').value,
            id: postId,
            userId: User.Id,
        }

        socket.send(JSON.stringify(formResponse));
    })
}