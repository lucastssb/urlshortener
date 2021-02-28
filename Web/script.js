async function onSubmitHandler(event) {
  event.preventDefault();

  const URL_TO_FETCH = "http://localhost:8080/create";

  const formData = document.getElementById("url-form").elements;
  const responseBox = document.getElementById('response-box');
  const responseText = document.getElementById('response');

  const url = formData["url"].value;

  const shortUrl = await fetch(URL_TO_FETCH, {
    method: "post",
    body: JSON.stringify({ url: url }),
  });

  const response = await shortUrl.json();

  if(response !== undefined) {
    responseBox.classList.remove('invisible');
    responseText.classList.remove('invisible');
    responseText.append(`http://localhost:8080${response.urlShorten}`);
  }

  console.log(response.urlShorten);
}
