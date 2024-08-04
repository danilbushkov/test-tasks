

async function sendFile() {
  const input = document.getElementById("file-input");
  const error = document.getElementById("import-error");

  let file = input.files[0];

  if(file != null) {
    let data = new FormData();
    data.append('file', file);

    let response = await fetch("/api/products/import", {
      method: "POST",
      body: data,
    });
    if(response.ok) {
      location.href = "/products";
    } else if(response.status === 422) {
      let json = await response.json();
      error.innerHTML = json.error.message;
    } else {
      error.innerHTML = "Произошла ошибка";
    }
  } else {
    error.innerHTML = "Выберите файл";
  }
}
