import SearchHome from "../../components/searchHome/SearchHome";
import "./home.css";
import Navbar from "../../components/navbar/Navbar.tsx";
import DownArrow from "../../assets/down-arrow.svg";
import BottomImage from "../../assets/home_bottom_img.png";
import { useEffect, useState } from "react";
import { Category, getCategories } from "../../networking/getCategories.ts";

export const Home = () => {
  const [categories, setCategories] = useState<Category[]>();
  const [searchShown, setSearchShown] = useState<boolean>(false);

  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
    const categories = await getCategories();
    setCategories(categories);
  };

  const searchClicked = () => {
    setSearchShown(true);
  };

  console.log(categories);
  if (!searchShown) {
    return (
      <>
        <div className="container">
          <Navbar />
          <section className="main-section">
            <div className="home-hero">
              <h1>Home,</h1>
              <h1>improvement,</h1>
              <h1>made easy</h1>
            </div>
            <SearchHome onSearchClick={searchClicked} />
            <div>
              <h4 className="h4-searchTip">
                Try searching for a <span>Plumber</span>, <span>Handyman</span>,{" "}
                <span>Landscaper</span> or <span>Electrician</span>
              </h4>
            </div>
          </section>
          <div className="bottom-arrow">
            <img src={DownArrow} alt="scroll down arrow" />
            <img src={BottomImage} alt="bottom image" />
          </div>
        </div>
      </>
    );
  } else {
    return <div>search data</div>;
  }
};

export default Home;
