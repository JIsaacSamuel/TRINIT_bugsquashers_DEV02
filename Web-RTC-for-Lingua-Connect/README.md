# Web-RTC feature for Lingua Connect 
Server-Side Setup:

The server is built using Express.js framework.
It sets up routes to handle requests to the root URL ("/") and dynamic room URLs ("/:room").
When a client connects to the root URL, it generates a unique room ID using UUID v4 and redirects the client to that room URL.
When a client connects to a room URL, it renders an EJS template named "room" and passes the room ID as a parameter to the template.
Socket.IO Integration:

Socket.IO is used for real-time communication between clients and the server.
When a client connects to the server, a Socket.IO connection is established.
Clients can join specific rooms using the "join-room" event, passing the room ID and user ID as parameters.
Upon joining a room, clients subscribe to room-specific events such as "user-connected" and "user-disconnected" to handle signaling for WebRTC connections.
Client-Side Setup:

The client-side code is written in JavaScript and runs in the browser.
It establishes a connection to the server using Socket.IO.
It initializes a PeerJS instance (using PeerJS library) to handle WebRTC peer-to-peer connections.
When the client's browser supports WebRTC, it prompts the user for access to their camera and microphone.
Upon granting access, it captures the user's video and audio streams using navigator.mediaDevices.getUserMedia.
It renders the user's video stream in a video element on the page.
It listens for incoming calls from other users (via the "call" event) and answers them by sending its own video stream.
It listens for "user-connected" events from the server and establishes WebRTC connections with newly connected users.
It listens for "user-disconnected" events from the server and closes WebRTC connections with disconnected users.
It dynamically adds video elements for each connected user's video stream and appends them to the HTML page.
In summary, the application allows users to create or join video conferencing rooms, where they can communicate with each other via real-time video and audio streams using WebRTC technology, with Socket.IO handling the signaling process for establishing peer-to-peer connections.
