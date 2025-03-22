import LogoutOutlinedIcon from '@mui/icons-material/LogoutOutlined';
import Link from 'next/link';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import './profile.css';
import { useWorker } from '@/app/_Context/WorkerContext';

function Profilepop() {
    const [err, setErr] = useState('');
    const [user, setUser] = useState({});
    const router = useRouter();
    const {portRef} = useWorker()

    useEffect(() => {
            const storedUser = JSON.parse(localStorage.getItem('user')) || {};
            setUser(storedUser);
    }, []);

    const handleLogout = async () => {
        try {
            const response = await fetch("/api/logout", {
                method: "POST",
                credentials: "include",
            });

            if (response.status === 200) {                
                localStorage.removeItem('user');
                portRef?.current?.postMessage({
                    kind: "close"
                })
                router.push('/login')
            } else {
                setErr("Error while logging out.");
            }
        } catch (error) {
            setErr("Network error or server is unreachable.");
        }
    };

    return (
        <div className='profile-container'>
            {err && <div className='profileerr'>{err}</div>}
            {user?.id ? (
                <>
                    <div className='profile-div'>
                        <Link href={`/profile/${user.id}`}>
                        <div className='profile'>
                            <img className='left-profile' src={user.image || "/default-profile.png"} alt='Profile' />
                            <h3 className='right-profile'>{user.first_name} {user.last_name}</h3>
                        </div>
                        </Link>
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
