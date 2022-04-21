describe('login', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
    })

    it('loginLogout', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
        cy.url().should('eq', 'http://localhost:8000/welcome/Items')
        cy.get('#root > div > section > div.ant-layout > header.ant-layout-header.ant-pro-fixed-header > div > div.ant-space.ant-space-horizontal.ant-space-align-center.right___3L8KG > div:nth-child(4)').click()
        cy.contains('退出登录').click()

    })

    it('incorrectTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('include', '/user/login')
        // cy.url().should('eq', 'http://localhost:8000/user/login')
    })

})