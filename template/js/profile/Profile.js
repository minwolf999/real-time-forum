import { Home } from "../home/home.js";
import { SetNewCSS } from "../setnewcss.js";
import { socket, pageName, User } from "../index.js";
import { Modification } from "./modification.js";
import { Post } from "../post/post.js";
import { Chat } from "../chat.js";

/**
 * manage the profile page
 * @param {boolean} filtered 
 */
export const Profile = async (filtered = false) => {
    pageName.pagename = 'profile'
    SetNewCSS(['/template/css/profile.css'])
    var data;

    if (!filtered) {
        data = await fetch('http://localhost:8080/getposts05842365165', {method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())
    } else {
        data = await fetch('http://localhost:8080/getcategories05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())
    }

    const notification = await fetch('http://localhost:8080/getnotificationQuantity05842365165', { method: 'POST', body: JSON.stringify([User.Id, User.Username, User.Age, User.Gender, User.FirstName, User.LastName, User.Email, User.Password]) }).then(async data => data.json())

    document.getElementById('content').innerHTML = `
    <div class="barre_menu">
        <button type="button" id='home' class="home">Home</button>

        <h2 class="text">Home Page</h2>

        <div class="messagerieDiv">
            <button type="button" id="messagerie" class="messagerie">Messagerie</button>
            <div id="notification" class="notification">${notification}</div>
        </div>
    </div>

    <div class="body">
    <div class="info">
        <img src="${User.ImagePath}">
        <div>
            <p><i>Name:</i></p>
            <p>${User.Username}</p>
        </div>
    
        <div>
            <p><i>Email:</i></p>
            <p>${User.Email}</p>
        </div>
    
        <div>
            <p><i>Registration Date:</i></p>
            <p>${User.RegistrationDate}</p>
        </div>
        <br>

        <button class="modification" id="modificationUserLink">Modification User's Information</button>
        </div>

        <div class="filtres_button">
            <p>Filter history</p>
            <div>
                <button name="filtre" value="myPost">My Posts</button>
                <button name="filtre" value="myComment">My Comments</button>
                <button name="filtre" value="myLikePost">My Likes Posts</button>
                <button name="filtre" value="myDislikePost">My Dislikes Posts</button>
            </div>
        </div>

        <div class="titled_post">
            <p class="titre">History</p>

            <div class="post" id='post'></div>
        </div>
    </div>
    `

    if (data) {
        data.Posts.forEach(post => {
            const previousPost = document.createElement('div')
            previousPost.className = 'previousPost'
    
            previousPost.addEventListener('click', () => {
                Post(post.Id)
            })
    
            document.getElementById('post').appendChild(previousPost)
    
            let categories = post.NameCategories.reduce((acc, current) => acc + ', ' + current, "")
            categories = categories.slice(2)
    
    
            const firstline = document.createElement('div')
            firstline.setAttribute('class', 'firstline')
            previousPost.appendChild(firstline)
    
            const img = document.createElement('img')
            img.src = post.CreatorImageProfile
            firstline.appendChild(img)
    
            const info = document.createElement('p')
            info.setAttribute('class', 'infoPost')
            info.textContent = `${post.NameCreator} - ${post.CreationDate}`
            firstline.appendChild(info)
    
            const likes_dislikes = document.createElement('p')
            likes_dislikes.setAttribute('class', 'like_dislikes')
            likes_dislikes.innerHTML = `&#x1F44D;${post.Likes.Quantity}   &#x1F44E; ${post.Dislikes.Quantity}`
            firstline.appendChild(likes_dislikes)
    
            const nameCategorie = document.createElement('p')
            nameCategorie.setAttribute('class', 'nameCategorie')
            nameCategorie.innerHTML = `
            <i>Categories:</i>
            ${categories}
            `
            previousPost.appendChild(nameCategorie)
    
            const br = document.createElement('br')
            previousPost.appendChild(br)
    
            const namePost = document.createElement('p')
            namePost.setAttribute('class', 'namePost')
            namePost.innerHTML = `
                <i>Name:</i>
                ${post.Name}
            `
            previousPost.appendChild(namePost)
    
            const descriptionPostSpan = document.createElement('span')
            descriptionPostSpan.innerHTML = '<i>Description:</i>'
            previousPost.appendChild(descriptionPostSpan)
    
            const descriptionPostPre = document.createElement('pre')
            descriptionPostPre.setAttribute('class', 'descriptionPost')
            descriptionPostPre.innerHTML = post.Description
            previousPost.appendChild(descriptionPostPre)
        });
    }

    document.getElementById('messagerie').addEventListener('click', () => {
        Chat()
    })

    document.getElementById('home').addEventListener('click', () => {
        Home()
    })

    document.getElementById('modificationUserLink').addEventListener('click', () => {
        Modification()
    })

    document.getElementsByName('filtre').forEach(button => {
        button.addEventListener('click', () => {
            let formResponse = {}

            formResponse[pageName.pagename] = {
                type: 'filter',
                filterType: button.value.toString(),
                userId: User.Id,
            }

            socket.send(JSON.stringify(formResponse));
        })
    })
}
