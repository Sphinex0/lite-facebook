import LogoutOutlinedIcon from '@mui/icons-material/LogoutOutlined';
import Link from 'next/link';
import { useState, useEffect, useRef } from 'react';
import { useRouter } from 'next/navigation';
import './profile.css';
import { FetchApi } from '@/app/helpers';
import { useContext } from 'react';
import { WorkerContext } from '@/app/_Context/WorkerContext';

function Profilepop({ setprofile }) {
    const [err, setErr] = useState('');
    const [user, setUser] = useState({});
    const router = useRouter();
    const refP = useRef(null)

    useEffect(() => {
        refP.current.focus()
        const storedUser = JSON.parse(localStorage.getItem('user')) || {};
        setUser(storedUser);
    }, []);

    const handleLogout = async () => {
        try {
            const response = await FetchApi("/api/logout", router, {
                method: "POST",
            });

            if (response.status === 200) {
                localStorage?.removeItem('user');
                router.push('/login')
            } else {
                setErr("Error while logging out.");
            }
        } catch (error) {
            console.log(error)
            setErr("Network error or server is unreachable.", error);
        }
    };

    return (
        <div tabIndex="0" className='profile-container' ref={refP} onBlur={() => setprofile(false)} >
            {err && <div className='profileerr'>{err}</div>}
            {user?.id ? (
                <>
                    <div className='profile-div'>
                        <div onClick={() => {
                            router.push(`/profile/${user?.id}`)
                            setprofile(false)
                        }}>
                            <div className='profile'>
                                <img src={`http://localhost:8080/public/pics/${user?.image || "default-profile.png"}`} alt='Profile'  className='left-profile' />
                                <h3 className='right-profile'>{user?.first_name} {user?.last_name}</h3>
                            </div>
                        </div>
                    </div>

                    <div onClick={handleLogout} className='profile-div'>
                        <div className='logout'>
                            <LogoutOutlinedIcon />
                        </div>
                        <h3>Logout</h3>
                    </div>

                </>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
}

export default Profilepop;
