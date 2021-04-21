
// self.addEventListener("fetch", function(event) {
//   event.respondWith(
//     fetch(event.request).catch(function() {
//       return new Response(
//       "I sit on a man’s back choking him and making him carry me, and yet assure myself and others that I am sorry for him and wish to lighten his load by all means possible…except by getting off his back."+
//       "Tolstoi"
//       );
//     })
//   );
// });
function isImage(fetchRequest) {
  return fetchRequest.method === "GET"
         && fetchRequest.destination === "image";
}

self.addEventListener("fetch", function(event) {
  console.log("Fetch request for:", event.request.url);
  if (event.request.url.includes("segato")) {
    console.log("Fetch request for:", event.request.url);
    event.respondWith(
      fetch("library.jpg"),
      new Response(
        "body {background: black!important;}",
        { headers: { "Content-Type": "text/css" }}
      )
    );
  }
  if (event.request.url) {
    event.respondWith(
      fetch(event.request)
        .then((response) => {
            if (response.ok) return response;
            if (isImage(event.request)) {
                return fetch("broken.png");
            }
        })
        .catch((err) => {
            if (isImage(event.request)) {
                console(err)
            }
        })
    );
  }
});