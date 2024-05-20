const image1 = document.getElementById("image1");
const image2 = document.getElementById("image2");
const image3 = document.getElementById("image3");

image1.style.opacity = 1;
image2.style.opacity = 0;
image3.style.opacity = 0;

setInterval(() => {
  image1.style.opacity = 0;
  image2.style.opacity = 1;
  image3.style.opacity = 0;

  setTimeout(() => {
    image1.style.opacity = 0;
    image2.style.opacity = 0;
    image3.style.opacity = 1;

    setTimeout(() => {
      image1.style.opacity = 1;
      image2.style.opacity = 0;
      image3.style.opacity = 0;
    }, 4000);
  }, 4000);
}, 12000);