"use client";

async function JoinGroup(id,creator,setIsAction,isAction) {
    console.log(id);
    console.log(creator);



    try {
        const response = await fetch('http://localhost:8080/api/invite/store', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({ "group_id": parseInt(id), "receiver": creator }),
        });

        if (response.ok) {
            if (isAction==="pending"){
                setIsAction("join")
            }else if (isAction==="join"){
                setIsAction("pending")
            }
            console.log('Group joined successfully');

        } else {
            console.error('Failed to join group');
        }
    } catch (error) {
        console.error('Error:', error);
    }



}

export default JoinGroup;
