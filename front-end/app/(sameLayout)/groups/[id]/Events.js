'use client'
import style from "./group.module.css";
import { Add, DisabledByDefault } from "@mui/icons-material";
import { useState, useEffect } from "react";

const Events = ({ groupID }) => {
    const [eventsData, setEventsData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');




    const CreateGroup = () => {



        const element = document.querySelector('#formId')
        element.style.display = "flex"
    }

    const SeeClick = () => {
        const element = document.querySelector('#formId')
        element.style.display = "none"
    }

    const handleSubmit = async (e) => {

            e.preventDefault();





            setTitle("")
            setDescription("")

            try {
                const response = await fetch("http://localhost:8080/api/Event/store", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    credentials: "include",
                    body:   JSON.stringify({ group_id: parseInt(groupID),Title:title,Description:description})
                });

                console.log(JSON.stringify({ group_id: parseInt(groupID) }));

                if (response.ok) {
                    const data = await response.json();
                    console.log(data);
                    setEventsData(data);
                    setLoading(false);
                    fetchEvents()
                } else {
                    setError("Failed to fetch events");
                    setLoading(false);
                }
            } catch (error) {
                setError("An error occurred while fetching events.");
                setLoading(false);
            }
        };

        
    





    useEffect(() => {

        const fetchEvents = async () => {
            try {
                const response = await fetch("http://localhost:8080/api/Events", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    credentials: "include",
                    body: JSON.stringify({ group_id: parseInt(groupID) }),
                });

                console.log(JSON.stringify({ group_id: parseInt(groupID) }));

                if (response.ok) {
                    const data = await response.json();
                    console.log(data);
                    setEventsData(data);
                    setLoading(false);
                } else {
                    setError("Failed to fetch events");
                    setLoading(false);
                }
            } catch (error) {
                setError("An error occurred while fetching events.");
                setLoading(false);
            }
        };

        fetchEvents();
    }, [groupID]);

    if (loading) {
        return <div>Loading...</div>;
    }
    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div>
            <div>
                <button className={style.create}
                    onClick={() => {
                        CreateGroup()
                    }}>
                    <Add />
                </button>
            </div>

            <div className={style.formclass} id="formId">
                <div onClick={() => { SeeClick() }}>
                    <DisabledByDefault />
                </div>
                <form method='POST' aria-multiselectable  onSubmit={handleSubmit}>
                    <label htmlFor='title' >title</label>
                    <input type='text' className={style.InputTitle} id='title' value={title} onChange={(e) => setTitle(e.target.value)} />
                    <label htmlFor='descriptopn'>descriptopn</label>
                    <input type='text' className={style.InputDescriptopn} id='descriptopn' value={description} onChange={(e) => setDescription(e.target.value)} />
                    <button className={style.InputButton} type='onSubmit'> Submit </button>
                </form>
            </div>




            <div>
                {eventsData.length === 0 ? (
                    <div>No events found.</div>
                ) : (
                    <ul>
                        {eventsData.map((event, index) => (
                            <div className='feed'>
                                <div className="user">
                                    <div className="ingo">
                                        <h3>{event.title} </h3>
                                        <small>{event.description}</small>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </ul>
                )}
            </div>
        </div>
    );
};

export default Events;
