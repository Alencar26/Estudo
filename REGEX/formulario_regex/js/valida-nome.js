export default function ehUmNome(campo) {
    const nome = campo.value
    /*
        ^ ->> Ancora indicando início do pattern
        $ ->> Ancora indicando final do pattern
        [A-Za-zÀ-ÿ -] ->> Indica que o pattern deve ser composto por letras maiúsculas, minúsculas, acentuação, espaço e ifem
        {3,30} ->> Indica que o pattern deve ter no mínimo 3 caracteres e no máximo 30 caracteres
        (?!(.)\1\1) ->> Negando as primeiras ocorrências iguail
    */
    const patternNome = /^(?!(.)\1\1)[A-Za-zÀ-ÿ -]{3,30}$/i.test(nome)
    console.log(patternNome);

    if (!patternNome) {

        campo.setCustomValidity('Não é um nome válido');
        console.log(`"${nome}" não é um nome válido`)

    }

    return nome
}