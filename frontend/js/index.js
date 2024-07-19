const url = 'https://beagle-mighty-terribly.ngrok-free.app/api/v1/tasks';

fetch(url, {
  method: 'GET',
  headers: {
    'ngrok-skip-browser-warning': 'anyvalue'
  }
})
.then(response => {
  if (!response.ok) {
    throw new Error('Network response was not ok ' + response.statusText);
  }
  return response.json();
})
.then(data => {

  const container = document.getElementById('js');
  
  data.forEach(task => {

    const titleDiv = document.createElement('div');
    titleDiv.textContent = `${task.title}`;
    container.appendChild(titleDiv);

})
.catch(error => {
  console.error('Error fetching JSON:', error);
})})


document.getElementById('form').addEventListener('submit', function(event) {
  event.preventDefault();  // Останавливает стандартное поведение отправки формы

  const formData = new FormData(this);
  const title = formData.get('input');

  fetch('https://beagle-mighty-terribly.ngrok-free.app/api/v1/createTask', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ title: title })
  })
  .then((response) => response.json())
  .then((data) => {
    fetch(url, {
      method: 'GET',
      headers: {
        'ngrok-skip-browser-warning': 'anyvalue'
      }
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok ' + response.statusText);
      }
      return response.json();
    })
    .then(data => {
    
      const container = document.getElementById('js');
      
      data.forEach(task => {
    
        const titleDiv = document.createElement('div');
        titleDiv.textContent = `${task.title}`;
        container.appendChild(titleDiv);
        
    })
    .catch(error => {
      console.error('Error fetching JSON:', error);
    })})
  })
  .catch((error) => {
    console.error('Error:', error);
  });
});


document.getElementById('delet').addEventListener('click', function() {
  fetch(url, {
    method: 'GET',
    headers: {
      'ngrok-skip-browser-warning': 'anyvalue'
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok ' + response.statusText);
    }
    return response.json();
  })
  .then(data => {
    const deletePromises = data.map(task => 
      fetch(`https://beagle-mighty-terribly.ngrok-free.app/api/v1/deleteTask/${task.id}`, {
        method: 'DELETE',
        headers: {
          'ngrok-skip-browser-warning': 'anyvalue'
        }
      })
    );

    return Promise.all(deletePromises);
  })
  .then(() => {
    const container = document.getElementById('js');
    container.innerHTML = ''; 
  })
  .catch(error => {
    console.error('Error:', error);
  });
});