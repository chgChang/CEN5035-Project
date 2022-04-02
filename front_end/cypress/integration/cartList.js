describe('cartList', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
        cy.get('#root > div > section > aside > div > div:nth-child(1) > ul > li:nth-child(2)').click()
    })
    it('proceedTest', function(){
        cy.get('.ant-pro-grid-content').get('button:contains("checkout")').click()
        cy.url().should('eq', 'http://localhost:8000/checkout')
    })
    // it('proceedTest', function(){
    //     cy.get('#root > div > section > div.ant-layout > main > div > div > div.ant-pro-grid-content > div > div > div > div > div.ant-card-body > div > div.ant-spin-nested-loading > div > ul > li:nth-child(1) > div.ant-input-number')
    //     cy.url().should('eq', 'http://localhost:8000/checkout')
    // })

})