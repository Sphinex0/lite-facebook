'use client'
import { FetchApi } from "@/app/helpers";
import style from "./group.module.css";
import { Add, DisabledByDefault } from "@mui/icons-material";
import EventsOptions from "./EventsOptions";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

const Events = ({ groupID }) => {
    const [eventsData, setEventsData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const redirect = useRouter()

    const CreateGroup = () => {
        const element = document.querySelector('#formId');
        element.style.display = "flex";
    }

    const SeeClick = () => {
        const element = document.querySelector('#formId');
        element.style.display = "none";
    }

    const handleSubmit = async (e) => {
        e.preventDefault();

        setTitle("");
        setDescription("");

        try {
            const response = await FetchApi("/api/Event/store",redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ group_id: parseInt(groupID), Title: title, Description: description })
            });

            if (!response.ok) {
                setError("An error occurred while submitting the event.");
                setLoading(false);
            } else {
                fetchEvents()
            }
        } catch (error) {
            setError("An error occurred while submitting the event.");
            setLoading(false);
        }
    };

    const fetchEvents = async () => {
        try {
            const response = await FetchApi("/api/Events",redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ group_id: parseInt(groupID) }),
            });


            if (response.ok) {
                const data = await response.json();
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

    useEffect(() => {
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
                <button className={style.create} onClick={CreateGroup}>
                    <Add />
                </button>
            </div>

            <div className={style.formclass} id="formId">
                <div onClick={SeeClick}>
                    <DisabledByDefault />
                </div>
                <form method='POST' aria-multiselectable onSubmit={handleSubmit}>
                    <label htmlFor='title'>Title</label>
                    <input type='text' className={style.InputTitle} id='title' value={title} onChange={(e) => setTitle(e.target.value)} />
                    <label htmlFor='description'>Description</label>
                    <input type='text' className={style.InputDescriptopn} id='description' value={description} onChange={(e) => setDescription(e.target.value)} />
                    <button className={style.InputButton} type='submit'>Submit</button>
                </form>
            </div>

            <div>
                {eventsData.length === 0 ? (
                    <div>No events found.</div>
                ) : (
                    <ul>
                        {eventsData.map((event, index) => (
                            <div className='feed' key={index}>
                                <div className={style.user}>
                                    <div className="info">
                                        <h3>{event.title}</h3>
                                        <small>{event.description}</small>
                                    </div>
                                    <div>
                                        <EventsOptions event_id={event.id}/>
                                        {/*  */}
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
