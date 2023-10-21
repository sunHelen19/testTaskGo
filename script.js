//    Получение через fetch список валют

let url =
  "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1";

let rowNumber = 1;

function fetchTableData(url) {
  fetch(url)
    .then((response) => response.json())
    .then((data) => {
      const table = document.getElementById("data");

      data.forEach((item) => {
        const row = table.insertRow();

        // Смена фона на синий для первых 5 валют
        if (rowNumber < 6) {
          row.style.backgroundColor = "blue";
          rowNumber += 1;
        }

        const idCell = row.insertCell();
        const symbolCell = row.insertCell();
        const nameCell = row.insertCell();
        idCell.textContent = item.id;
        symbolCell.textContent = item.symbol;
        nameCell.textContent = item.name;

        //  Смена фона на зеленый для валюты, где "symbol" = "usdt"
        if (item.symbol === "usdt") {
          changeBackgroundColorCell("green", row);
        }
      });
    })
    .catch((error) => console.error("Ошибка:", error));
}

fetchTableData(url);

function changeBackgroundColorCell(color, node) {
  node.style.backgroundColor = color;
}
