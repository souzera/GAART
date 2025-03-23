export interface IconButtonInterface{
    icon: string //url do icone
    label: string //texto que aparece ao passar o mouse
    onClick?: () => void //função que será executada ao clicar no botão
}