// Language toggle functionality
document.getElementById('toggleLanguage').addEventListener('click', function() {
    const englishContent = document.getElementById('english-content');
    const amharicContent = document.getElementById('amharic-content');
    const currentLanguage = document.getElementById('currentLanguage');

    if (englishContent.classList.contains('hidden')) {
        englishContent.classList.remove('hidden');
        amharicContent.classList.add('hidden');
        currentLanguage.textContent = 'English';
    } else {
        englishContent.classList.add('hidden');
        amharicContent.classList.remove('hidden');
        currentLanguage.textContent = 'አማርኛ';
    }
});

// FAQ toggle functionality
function toggleFAQ(id) {
    const content = document.getElementById(`faq-content-${id}`);
    const icon = document.getElementById(`faq-icon-${id}`);

    content.classList.toggle('hidden');
    icon.classList.toggle('rotate-180');
}

// Ethiopia regions data with coordinates and sample woredas
const ethiopiaRegions = {
    'Addis Ababa': {
        coordinates: { x: 50, y: 50, width: 30, height: 30 },
        woredas: ['Arada', 'Kirkos', 'Kolfe Keranio', 'Lideta', 'Nifas Silk-Lafto', 'Yeka']
    },
    'Afar': {
        coordinates: { x: 150, y: 50, width: 80, height: 60 },
        woredas: ['Asayita', 'Awash', 'Dubti', 'Gewane', 'Mille']
    },
    'Amhara': {
        coordinates: { x: 100, y: 100, width: 100, height: 100 },
        woredas: ['Bahir Dar', 'Debre Markos', 'Debre Tabor', 'Dessie', 'Gondar', 'Kombolcha']
    },
    'Benishangul-Gumuz': {
        coordinates: { x: 50, y: 150, width: 80, height: 60 },
        woredas: ['Asosa', 'Bambasi', 'Dangur', 'Mao-Komo', 'Sherkole']
    },
    'Dire Dawa': {
        coordinates: { x: 200, y: 100, width: 30, height: 30 },
        woredas: ['Gurgura', 'Hara']
    },
    'Gambela': {
        coordinates: { x: 50, y: 200, width: 60, height: 40 },
        woredas: ['Abobo', 'Gambela', 'Gog', 'Itang']
    },
    'Harari': {
        coordinates: { x: 180, y: 80, width: 20, height: 20 },
        woredas: ['Amir-Nur', 'Abadir', 'Shenkor']
    },
    'Oromia': {
        coordinates: { x: 120, y: 120, width: 150, height: 150 },
        woredas: ['Adama', 'Ambo', 'Bishoftu', 'Jimma', 'Nekemte', 'Shashamane']
    },
    'Somali': {
        coordinates: { x: 250, y: 100, width: 120, height: 100 },
        woredas: ['Degehabur', 'Gode', 'Jijiga', 'Kebri Beyah', 'Warder']
    },
    'SNNPR': {
        coordinates: { x: 120, y: 200, width: 120, height: 100 },
        woredas: ['Arba Minch', 'Awasa', 'Dilla', 'Sodo', 'Wolayta']
    },
    'Tigray': {
        coordinates: { x: 100, y: 50, width: 80, height: 60 },
        woredas: ['Adigrat', 'Axum', 'Mekelle', 'Shire', 'Wukro']
    }
};

// Initialize map functionality for both languages
function initializeMap(language) {
    const mapOverlay = document.getElementById(`map-overlay-${language}`);
    const mapImg = document.getElementById(`ethiopia-map-${language}`);

    // Clear previous highlights
    mapOverlay.innerHTML = '';

    // Create region highlights
    for (const [region, data] of Object.entries(ethiopiaRegions)) {
        const highlight = document.createElement('div');
        highlight.className = 'region-highlight';
        highlight.style.left = `${data.coordinates.x}px`;
        highlight.style.top = `${data.coordinates.y}px`;
        highlight.style.width = `${data.coordinates.width}px`;
        highlight.style.height = `${data.coordinates.height}px`;

        highlight.addEventListener('click', () => {
            showWoredaList(region, data.woredas, language);
            document.getElementById(`selected-region-${language}`).textContent = region;
            document.getElementById(`selected-region-${language}`).classList.remove('hidden');
            document.getElementById(`region-${language}`).value = region;
        });

        mapOverlay.appendChild(highlight);
    }
}

// Show woreda list for selected region
function showWoredaList(region, woredas, language) {
    const woredaList = document.getElementById(`woreda-list-${language}`);
    const woredaItems = document.getElementById(`woreda-items-${language}`);
    const regionTitle = document.getElementById(`selected-region-title-${language}`);

    // Set region title
    regionTitle.textContent = language === 'en' ? `Select Woreda in ${region}` : `ወረዳ ይምረጡ በ${region}`;

    // Clear previous woredas
    woredaItems.innerHTML = '';

    // Add woredas to list
    woredas.forEach(woreda => {
        const item = document.createElement('div');
        item.className = 'woreda-item';
        item.textContent = woreda;

        item.addEventListener('click', () => {
            document.getElementById(`woreda-${language}`).value = woreda;
            hideWoredaList(language);
        });

        woredaItems.appendChild(item);
    });

    // Show woreda list
    woredaList.classList.add('active');
}

// Hide woreda list
function hideWoredaList(language) {
    document.getElementById(`woreda-list-${language}`).classList.remove('active');
}

// Zoom map functionality
function zoomMap(scale, language) {
    const mapImg = document.getElementById(`ethiopia-map-${language}`);
    const currentWidth = parseInt(mapImg.style.width) || 100;
    const newWidth = currentWidth * scale;

    // Limit zoom levels
    if (newWidth < 80 || newWidth > 150) return;

    mapImg.style.width = `${newWidth}%`;
  mapImg.style.height = `${newWidth}%`;

    // Reinitialize map with new coordinates
    initializeMap(language);
}

// Initialize maps when page loads
window.addEventListener('DOMContentLoaded', () => {
    initializeMap('en');
    initializeMap('am');
});
