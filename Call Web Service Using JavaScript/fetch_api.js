//1st approach

fetch('https://jsonplaceholder.typicode.com/users')
    .then(response => {
        return response.json()
    }).then(json => {
        //console.log(json);
    })

//2nd approach

async function getUsers() {
    let response = await fetch('https://jsonplaceholder.typicode.com/users');
    let data = await response.json();
    return data;
}

getUsers().then(response=>{
    console.log(response);
})
