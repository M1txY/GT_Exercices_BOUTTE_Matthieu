const carousel = document.querySelector('.carousel');
const carouselContainer = carousel.querySelector('.carousel-container');
const carouselItems = carousel.querySelectorAll('.carousel-item');
const prevButton = carousel.querySelector('.carousel-prev');
const nextButton = carousel.querySelector('.carousel-next');

let currentIndex = 0;
let itemWidth = carouselItems[0].offsetWidth;

function showCurrentIndex() {
  carouselContainer.style.transform = `translateX(${-currentIndex * itemWidth}px)`;
}

prevButton.addEventListener('click', () => {
  currentIndex = currentIndex > 0 ? currentIndex - 1 : carouselItems.length - 1;
  showCurrentIndex();
});

nextButton.addEventListener('click', () => {
  currentIndex = currentIndex < carouselItems.length - 1 ? currentIndex + 1 : 0;
  showCurrentIndex();
});

window.addEventListener('resize', () => {
  itemWidth = carouselItems[0].offsetWidth;
  showCurrentIndex();
});
