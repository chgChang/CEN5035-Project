describe('register', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-tabs-tab.ant-tabs-tab-active').click()
    })
    it('correctTest', function(){
        // cy.get('button:contains("Login")').click()
        // cy.url().should('eq', 'http://localhost:8000/')
    })

    // it('incorrectTest', function(){
    //     cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
    //     cy.get('.ant-pro-form-login-main').get('#password').type('1234')
    //     cy.get('button:contains("Login")').click()
    //     cy.url().should('include', '/user/login')
    // })
})