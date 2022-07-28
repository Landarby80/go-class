import React from 'react'
import { useNavigate} from "react-router-dom";

import 'bootstrap/dist/css/bootstrap.min.css'

function WelcomePage() {
	const navigate = useNavigate();

  const usernamePage = () => {
    // 👇️ navigate to /contacts
    navigate('/createUsername');
  };

	return (
		<>

			<div className="jumbotron text-center">
				<h1 className="display-4">Welcome to chat app</h1>
				<p className="lead">
					It is an app where you can chat with friends.
				</p>

				<hr className="my-4" />
				<p>
					You can create channels or chat directly with your friends
				</p>

				<p className="lead">
					<button className="btn btn-success"
						onClick={usernamePage}>create an username
					</button>
				</p>

			</div>
		</>
	)
}



 export default WelcomePage;