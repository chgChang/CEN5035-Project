describe('cartList', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
    })

    it('addDelete', function () {
        cy.get('#root > div > section > div > main > div > div.ant-pro-grid-content > div > div > div > div > div > div > div > div > div:nth-child(1) > div')
          .contains("Add to Cart").click()
        cy.get('#root > div > section > aside > div > div:nth-child(1) > ul > li:nth-child(2)').click()
        cy.contains('More').click()
        cy.contains('Delete').click()
        cy.contains('Ok').click()
        cy.url().should('include', 'cart-list')
        cy.wait(5000)
        cy.contains('No Data')

        cy.get('.ant-pro-grid-content').get('button:contains("checkout")').click()
        cy.url().should('eq', 'http://localhost:8000/checkout')
    })
})