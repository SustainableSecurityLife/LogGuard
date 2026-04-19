const app = document.getElementById("app")!

fetch("http://localhost:8080/analyze")
  .then(res => res.json())
  .then(data => {
    data.forEach((item: any) => {
      const p = document.createElement("p")
      p.textContent = `${item.ip} : ${item.count}`
      app.appendChild(p)
    })
  })
  .catch(err => {
    console.error(err)
    app.textContent = "Error connecting to API"
  })