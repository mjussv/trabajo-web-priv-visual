const API_URL = '/looks'; // O '/prendas' según el endpoint que maneje tu servidor Go

const formPrenda = document.getElementById('form-prenda');
const listaPrendas = document.getElementById('lista-prendas');

// 1. Cargar outfits al iniciar la página (GET)
document.addEventListener('DOMContentLoaded', obtenerOutfits);

function obtenerOutfits() {
  fetch(API_URL)
    .then(res => res.json())
    .then(data => {
      listaPrendas.innerHTML = '';
      if (!data || data.length === 0) {
        listaPrendas.innerHTML = '<li style="text-align:center; color:#777;">No hay outfits guardados todavía.</li>';
        return;
      }

      data.forEach(item => {
        const li = document.createElement('li');
        li.className = 'item-prenda';
        li.innerHTML = `
          <div class="item-info">
            <strong>${item.nombre}</strong>
            <p>${item.descripcion || ''}</p>
            <span class="tags">${item.ocasion || ''} | ${item.temporada || ''} | ${item.estilo || ''}</span>
          </div>
          <button class="btn-eliminar" data-id="${item.id}">Eliminar</button>
        `;
        listaPrendas.appendChild(li);
      });
    })
    .catch(err => console.error('Error al obtener outfits:', err));
}

// 2. Crear outfit enviando todos los campos (POST)
formPrenda.addEventListener('submit', (e) => {
  e.preventDefault();

  const nuevoOutfit = {
    usuario_id: parseInt(document.getElementById('usuario_id').value),
    nombre: document.getElementById('nombre').value,
    descripcion: document.getElementById('descripcion').value,
    ocasion: document.getElementById('ocasion').value,
    temporada: document.getElementById('temporada').value,
    url_imagen: document.getElementById('url_imagen').value,
    estilo: document.getElementById('estilo').value
  };

  fetch(API_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(nuevoOutfit)
  })
  .then(res => {
    if (res.ok) {
      formPrenda.reset();
      obtenerOutfits();
    } else {
      alert('Error al guardar el outfit.');
    }
  })
  .catch(err => console.error('Error al crear:', err));
});

// 3. Borrado con delegación de eventos (DELETE)
listaPrendas.addEventListener('click', (e) => {
  if (e.target.classList.contains('btn-eliminar')) {
    const id = e.target.getAttribute('data-id');
    fetch(`${API_URL}/${id}`, { method: 'DELETE' })
      .then(res => {
        if (res.ok) obtenerOutfits();
      })
      .catch(err => console.error('Error al eliminar:', err));
  }
});
