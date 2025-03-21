import LogoutOutlinedIcon from '@mui/icons-material/LogoutOutlined';
import Link from 'next/link';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import './profile.css';

function Profilepop() {
    const [err, setErr] = useState('');
    const [user, setUser] = useState({});
    const router = useRouter();

    useEffect(() => {
            const storedUser = JSON.parse(localStorage.getItem('user')) || {};
            setUser(storedUser);
    }, []);

    const handleLogout = async () => {
        try {
            const response = await fetch("http://localhost:8080/api/logout", {
                method: "POST",
                credentials: "include",
            });

            if (response.status === 200) {
                console.log("yesssss");
                
                localStorage.removeItem('user');
                // Try using window.location.href for redirection
                window.location.href = '/login';
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
                            <h3>{user.first_name} {user.last_name}</h3>
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
