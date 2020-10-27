import React from "react"
import { styled } from "@material-ui/core/styles"
import { Link } from "react-router-dom"

const StyledLink = styled(Link)(({ theme }) => ({
  color: theme.palette.primary.main,
  textDecoration: "none",
  fontSize: "18px",
}))

const List = styled("ul")({
  display: "flex",
  justifyContent: "flex-end",
  marginLeft: "auto",
  listStyle: "none",
  alignItems: "center",
})

const ListItem = styled("li")(({ theme }) => ({
  marginRight: theme.spacing(3),
  "&:last-of-type": {
    marginRight: 0,
  },
}))

const TopNav = () => (
  <nav>
    <List>
      <ListItem>
        <StyledLink to="/about">About</StyledLink>
      </ListItem>
    </List>
  </nav>
)

export default TopNav
