import { User, pageName, socket } from "../index.js";

/**
 * manage the display comment zone
 * 
 * @param {Array.Object} data 
 * @param {string} PostId 
 */
export const Comments = (data, PostId) => {
    data.forEach(comment => {
        const div = document.createElement('div')
        div.className = 'previousComment'

        div.innerHTML = `
            <input type="text" value="${comment.Id}" name="id" hidden>
            <div class="firstline">
                <img class="img" src="${comment.CreatorImagePath}">
                <p class="info">${comment.NameCreator} - ${comment.CreationDate}</p>
                <div class="like_dislikes">
                    <button type="submit" name="${comment.Id}actionComment" value="like">&#x1F44D; ${comment.Likes.Quantity}</button>
                    <button type="submit" name="${comment.Id}actionComment" value="dislike">&#x1F44E; ${comment.Dislikes.Quantity}</button>
                </div>
            </div>

            <span><i>Message:</i></span>
            <pre class="message">${comment.Text}</pre>
        `
        
        document.getElementById('bodyCenter').appendChild(div)

        document.getElementsByName(`${comment.Id}actionComment`).forEach(button => {
            button.addEventListener('click', () => {
                let formResponse = {}
    
                formResponse[pageName.pagename] = {
                    type: 'LikeComment',
                    value: button.value,
                    id: comment.Id,
                    postId: PostId,
                    userId: User.Id,
                }
                
                socket.send(JSON.stringify(formResponse));
            })
        })
        
    });
}