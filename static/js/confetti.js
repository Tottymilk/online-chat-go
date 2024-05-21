$(document).ready(function() {
    function createConfettiPiece() {
        var confettiPiece = $("<div class='confetti-piece'></div>");
        var x = Math.random() * $(window).width()-100;
        var y = Math.random() * $(window).height()-200;
        confettiPiece.css({
        left: x,
        top: y
        });
        return confettiPiece;
    }

    function animateConfettiPiece(confettiPiece) {
        var angle = Math.random() * 2 * Math.PI;
        var speed = Math.random() * 30 + 5;
        var x = Math.cos(angle) * speed;
        var y = Math.sin(angle) * speed;
        confettiPiece.animate({
        left: "+=" + x,
        top: "+=" + y
        }, 5000, function() {
        confettiPiece.remove();
        });
    }

    setInterval(function() { 
        var confettiPiece = createConfettiPiece();
        $("#confetti-container").append(confettiPiece);
        animateConfettiPiece(confettiPiece);
    }, 10);
});