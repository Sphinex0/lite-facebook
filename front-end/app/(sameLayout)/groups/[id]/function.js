"use client";

import { FetchApi } from "@/app/helpers";

async function JoinGroup(id,creator,setIsAction,isAction,redirect) {
    console.log(id);
    console.log(creator);



    try {
        const response = await FetchApi('/api/invite/store',redirect, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
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
