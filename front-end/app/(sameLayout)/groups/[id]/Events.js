const Events = async ({ groupID }) => {


    try {
        const response = await fetch("http://localhost:8080/api/Events", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: "include",
            body: JSON.stringify({ group_id: parseInt(groupID) }),

        })
        console.log(JSON.stringify({ group_id: parseInt(groupID) }));

        if (response.ok) {
            const EventsData = await response.json()
            console.log(EventsData);
        }


    } catch (error) {
        console.log(error)
    }

    
}


export default Events