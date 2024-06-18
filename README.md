# real-time-forum

## Description
This project aims to create a web forum allowing users to communicate, associate categories to posts, like/dislike posts and comments, and filter posts. The forum will be implemented using SQLite for data storage and will include features such as authentication, communication, likes/dislikes, filtering, chating.

## Authors
- Authors: Antoine Marvin

## Technologies Used
- SQLite
- Go (Golang)
- Js (Javascript) with websocket

## Features
1. **Authentication**:
   - Users can register with unique email addresses and usernames.
   - Encrypted password storage.

2. **Communication**:
   - Registered users can create posts and comments.
   - Posts can be associated with one or more categories.
   - Posts and comments are visible to all users.

3. **Likes and Dislikes**:
   - Registered users can like or dislike posts and comments.
   - Like and dislike counts are visible to all users.

4. **Filtering**:
   - Users can filter posts by categories, created posts, and liked posts.
   - Filtered categories act as subforums.

5. **Chat**:
   - Users can see if the other users are connected or not
   - Users can chat between them in real time