let url='https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1';



function fetchTableData(url) {
  fetch(url) 
    .then(response => response.json())
    .then(data => {
      const table = document.getElementById('data');
 
      data.forEach(item => {
        const row = table.insertRow();
        const idCell = row.insertCell();
        const symbolCell=row.insertCell();
        const nameCell = row.insertCell();
        idCell.textContent = item.id;
        symbolCell.textContent = item.symbol;
        nameCell.textContent=item.name;
      });
    })
    .catch(error => console.error('Ошибка:', error));
}

fetchTableData(url)
