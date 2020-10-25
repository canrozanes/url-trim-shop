import React from "react"
import { styled } from "@material-ui/core/styles"
import Typography from "@material-ui/core/Typography"
import Container from "@material-ui/core/Container"
import AppBar from "@material-ui/core/AppBar"
import Toolbar from "@material-ui/core/Toolbar"
import TopNav from "./topnav"
import { Link } from "react-router-dom"

const HeaderLogo = styled(Link)(({ theme }) => ({
  color: theme.palette.primary.main,
  textDecoration: "none",
}))

const StyledAppBar = styled(AppBar)({
  justifyContent: "space-between",
  background: "white",
  position: "static",
})

const StyledToolbar = styled(Toolbar)({
  display: "flex",
  justifyContent: "space-between",
})

const Header = () => (
  <StyledAppBar>
    <Container maxWidth="xl">
      <StyledToolbar>
        <Typography variant="h4">
          <HeaderLogo to="/" color="primary">
            Trimmer.ly
          </HeaderLogo>
        </Typography>
        <TopNav />
      </StyledToolbar>
    </Container>
  </StyledAppBar>
)

export default Header
