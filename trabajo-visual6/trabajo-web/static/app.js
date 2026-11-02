const API_LOOKS = '/looks';
const API_USUARIOS = '/usuarios';
const API_ME_GUSTA = '/me_gusta_looks';

let usuarioActual = null;

// Elementos del DOM
const vistaAuth = document.getElementById('vista-auth');
const vistaApp = document.getElementById('vista-app');
const formLogin = document.getElementById('form-login');
const formRegistro = document.getElementById('form-registro');
const mensajeAuth = document.getElementById('mensaje-auth');
const bienvenidaUsuario = document.getElementById('bienvenida-usuario');
const btnLogout = document.getElementById('btn-logout');

const formOutfit = document.getElementById('form-outfit');
const contenedorOutfits = document.getElementById('contenedor-outfits');

// 1. LOGIN CON EMAIL Y CONTRASEÑA
formLogin.addEventListener('submit', (e) => {
  e.preventDefault();
  mensajeAuth.textContent = '';
  const emailInput = document.getElementById('auth-email').value;
  const passInput = document.getElementById('auth-pass').value;

  fetch(API_USUARIOS)
    .then(res => res.json())
    .then(usuarios => {
      // Buscamos si existe el usuario con ese email
      const usuarioEncontrado = usuarios.find(u => u.email === emailInput);

      if (usuarioEncontrado) {
        usuarioActual = usuarioEncontrado;
        iniciarSesionEnVista(usuarioActual);
      } else {
        mensajeAuth.textContent = '⚠️ Usuario no encontrado. Podés registrarte abajo:';
        formRegistro.classList.remove('hidden');
      }
    })
    .catch(err => {
      console.error(err);
      mensajeAuth.textContent = '❌ Error al conectar con el servidor.';
    });
});

// 2. REGISTRO DE USUARIO
formRegistro.addEventListener('submit', (e) => {
  e.preventDefault();
  const nuevoUsuario = {
    nombre: document.getElementById('reg-nombre').value,
    email: document.getElementById('reg-email').value
  };

  fetch(API_USUARIOS, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(nuevoUsuario)
  })
  .then(res => {
    if (!res.ok) throw new Error('Error al registrar');
    return res.json();
  })
  .then(user => {
    usuarioActual = user;
    alert(`¡Cuenta creada con éxito! Bienvenida/o ${user.nombre}`);
    iniciarSesionEnVista(user);
  })
  .catch(err => console.error(err));
});

function iniciarSesionEnVista(usuario) {
  vistaAuth.classList.add('hidden');
  vistaApp.classList.remove('hidden');
  bienvenidaUsuario.textContent = `¡Hola, ${usuario.nombre}! ✨`;
  obtenerOutfits();
}

btnLogout.addEventListener('click', () => {
  usuarioActual = null;
  vistaApp.classList.add('hidden');
  vistaAuth.classList.remove('hidden');
  formRegistro.classList.add('hidden');
  formLogin.reset();
  formRegistro.reset();
});

// 3. OBTENER OUTFITS Y MOSTRAR FOTOS (GET)
function obtenerOutfits() {
  fetch(API_LOOKS)
    .then(res => res.json())
    .then(data => {
      contenedorOutfits.innerHTML = '';

      if (!data || data.length === 0) {
        contenedorOutfits.innerHTML = '<p class="empty-msg">No hay outfits cargados en tu guardarropa.</p>';
        return;
      }

      data.forEach(look => {
        const card = document.createElement('article');
        card.className = 'outfit-card';
        
        // Foto por defecto si la URL viene vacía
        const fotoUrl = look.url_imagen || 'https://via.placeholder.com/300x400?text=Sin+Foto';

        card.innerHTML = `
          <div class="img-container">
            <img src="${fotoUrl}" alt="${look.nombre}" onerror="this.src='https://via.placeholder.com/300x400?text=Imagen+No+Disponible'">
          </div>
          <div class="card-body">
            <h3>${look.nombre}</h3>
            <p class="desc">${look.descripcion || ''}</p>
            <div class="details">
              <span>🏷️ ${look.ocasion || 'General'}</span>
              <span>✨ ${look.estilo || 'Urbano'}</span>
            </div>
            <div class="card-actions">
              <button class="btn-like" data-id="${look.id}">❤️ Me gusta <span class="likes-count">${look.likes || 0}</span></button>
              <button class="btn-eliminar" data-id="${look.id}">🗑️ Borrar</button>
            </div>
          </div>
        `;
        contenedorOutfits.appendChild(card);
      });
    })
    .catch(err => console.error('Error al obtener outfits:', err));
}

// 4. CREAR OUTFIT CON FOTO (POST)
formOutfit.addEventListener('submit', (e) => {
  e.preventDefault();

  const nuevoLook = {
    usuario_id: parseInt(usuarioActual.id),
    nombre: document.getElementById('nombre').value,
    descripcion: document.getElementById('descripcion').value,
    ocasion: document.getElementById('ocasion').value,
    temporada: document.getElementById('temporada').value,
    url_imagen: document.getElementById('url_imagen').value,
    estilo: document.getElementById('estilo').value
  };

  fetch(API_LOOKS, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(nuevoLook)
  })
  .then(res => {
    if (res.ok) {
      formOutfit.reset();
      obtenerOutfits();
    } else {
      alert('Error al guardar el outfit.');
    }
  });
});

// 5. DELEGACIÓN DE EVENTOS (ME GUSTA & BORRADO)
contenedorOutfits.addEventListener('click', (e) => {
  const lookId = e.target.getAttribute('data-id') || e.target.parentElement.getAttribute('data-id');

  // Acción: DAR ME GUSTA
  if (e.target.classList.contains('btn-like') || e.target.parentElement.classList.contains('btn-like')) {
    fetch(API_ME_GUSTA, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        usuario_id: parseInt(usuarioActual.id),
        look_id: parseInt(lookId)
      })
    })
    .then(res => {
      // Refrescamos o aumentamos visualmente
      obtenerOutfits();
    })
    .catch(err => console.error('Error al dar Me Gusta:', err));
  }

  // Acción: ELIMINAR OUTFIT
  if (e.target.classList.contains('btn-eliminar')) {
    fetch(`${API_LOOKS}/${lookId}`, { method: 'DELETE' })
      .then(res => {
        if (res.ok) obtenerOutfits();
      });
  }
});
