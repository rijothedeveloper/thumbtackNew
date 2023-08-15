import SearchHome from "../../components/searchHome/SearchHome";
import "./home.css";

export const Home = () => {
  return (
    <>
      <div className="container">
        <section className="main-section">
          <div className="home-hero">
            <h1>Home,</h1>
            <h1>improvement,</h1>
            <h1>made easy</h1>
          </div>
          <SearchHome />
        </section>
      </div>
    </>
  );
};

export default Home;
