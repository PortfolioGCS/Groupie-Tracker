background=document.querySelector('.background')
bg=document.querySelector('.bg')
nom=document.querySelector('.nom')
user=document.querySelector('.user')
divbutton=document.querySelector('.divbutton')
divbutton2=document.querySelector('.divbutton2')

divbutton.onclick = function flou() {
    nom.style.filter = 'blur(0px)';
    user.style.filter = 'blur(0px)';
    button.style.fontSize = '0px'
    button.style.height = '0px';
    button.style.width = '0px';
    setTimeout(disappear, 250)
    setTimeout(appear, 270)
  }

  
  function disappear() {
    button.style.display = 'none';
    
  }

function appear(){
  button2.style.visibility = 'visible';
  button2.style.fontSize = '3.75vh';
  button2.style.height = '116px';
  button2.style.width = '380px';
}
  

  console.log("OK")