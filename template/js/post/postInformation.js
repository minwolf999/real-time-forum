import { User, pageName, socket } from "../index.js"


/**
 * manage the display of the information of the post
 * 
 * @param {Array.Object} data 
 */
export const postInformation = (data) => {
    let categories = data.NameCategories.reduce((acc, current) => acc + ', ' + current, "")
    categories = categories.slice(2)

    document.getElementById("bodyLeft").innerHTML = `
        <div class="firstline">
            <div class="information">
                <img src="${data.CreatorImageProfile}">
                <p class="info">${data.NameCreator} - ${data.CreationDate}</p>
            </div>

            <div class="like_dislike" id="form">
                <button type="submit" name="action" value="like">&#x1F44D; ${data.Likes.Quantity}</button>
                <button type="submit" name="action" value="dislike">&#x1F44E; ${data.Dislikes.Quantity}</button>
            </div>
        </div>
        <br>

        <span><i>Description of the post:</i></span>
        <pre class="descriptionPost">${data.Description}</pre>
        <br>

        <p class="nameCategorie"><i>Categories of the post:</i>
            ${categories}
        </p><br>
    `

    document.getElementsByName("action").forEach(button => {
        button.addEventListener('click', () => {
            let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'likePost',
                value: button.value,
                id: data.Id,
                userId: User.Id,
            }

            socket.send(JSON.stringify(formResponse));
        })
    })
}