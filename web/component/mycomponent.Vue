<template>
  <div class="choose-date-range-components">
    <div class="opt-btns">
      <el-link @click="determine">确定</el-link>
      <el-link @click="cancel">取消</el-link>
    </div>
    <div class="tow-month">
      <div class="month-bar">
        <div class="header">
          <div class="pre-month">
            <i class="arrow el-icon-caret-left" @click="goPreMonth"></i>
          </div>
          <div class="month-date">{{curLeftYear}}年{{curLeftMonth}}月</div>
          <div class="next-month"></div>
        </div>
        <div class="weekdays">
          <div class="weekday">日</div>
          <div class="weekday">一</div>
          <div class="weekday">二</div>
          <div class="weekday">三</div>
          <div class="weekday">四</div>
          <div class="weekday">五</div>
          <div class="weekday">六</div>
        </div>
        <div class="days">
          <div class="day"
               @click="chooseDay(curLeftYear, curLeftMonth, day)"
               @mouseover="preChoose(curLeftYear, curLeftMonth, day)"
               @mouseout="preNotChoose(curLeftYear, curLeftMonth, day)"
               :class="{
                 'cannot-choose': judgeDisable(curLeftYear, curLeftMonth, day),
                 'in-range': judgeInRange(curLeftYear, curLeftMonth, day),
                 'in-pre-range': judgeInPreRange(curLeftYear, curLeftMonth, day),
                 'is-start': judgeStart(curLeftYear, curLeftMonth, day),
                 'is-end': judgeEnd(curLeftYear, curLeftMonth, day),
                 'is-today': judgeToday(curLeftYear, curLeftMonth, day)
               }"
               v-for="(day, index) in curLeftDays" :key="index">
            <div :class="{'day-inner': true,'chosen-day': judgeChosen(curLeftYear, curLeftMonth, day)}">
              {{day}}
            </div>
          </div>
        </div>
      </div>

      <div class="month-bar">
        <div class="header">
          <div class="pre-month"></div>
          <div class="month-date">{{curRightYear}}年{{curRightMonth}}月</div>
          <div class="next-month">
            <i class="arrow el-icon-caret-right" @click="goNextMonth"></i>
          </div>
        </div>
        <div class="weekdays">
          <div class="weekday">日</div>
          <div class="weekday">一</div>
          <div class="weekday">二</div>
          <div class="weekday">三</div>
          <div class="weekday">四</div>
          <div class="weekday">五</div>
          <div class="weekday">六</div>
        </div>
        <div class="days">
          <div class="day"
               @click="chooseDay(curRightYear, curRightMonth, day)"
               @mouseover="preChoose(curRightYear, curRightMonth, day)"
               @mouseout="preNotChoose(curRightYear, curRightMonth, day)"
               :class="{
                 'cannot-choose': judgeDisable(curRightYear, curRightMonth, day),
                 'in-range': judgeInRange(curRightYear, curRightMonth, day),
                 'in-pre-range': judgeInPreRange(curRightYear, curRightMonth, day),
                 'is-start': judgeStart(curRightYear, curRightMonth, day),
                 'is-end': judgeEnd(curRightYear, curRightMonth, day),
                 'is-today': judgeToday(curRightYear, curRightMonth, day)
               }"
               v-for="(day, index) in curRightDays" :key="index">
            <div :class="{'chosen-day': judgeChosen(curRightYear, curRightMonth, day)}">
              {{day}}
            </div>
          </div>
          <!--        <div class="day" v-for="(day, index) in curRightDays" :key="index">-->
          <!--          {{day}}-->
          <!--        </div>-->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ChooseDateRange",
  props: ['monthNum', 'startDateStr', 'endDateStr'],
  data () {
    return {
      leftDate: {},
      rightDate: {},
      chooseStartDateStr: '',
      chooseEndDateStr: '',
      preChooseDateStr: '',
      curLeftYear: '',
      curLeftMonth: '',
      curRightYear: '',
      curRightMonth: '',
      curLeftDays: [],
      curRightDays: []
    }
  },
  methods: {
    addMonth (date, monthNum = this.monthNum) {
      let year = date.getFullYear()
      let month = date.getMonth() + 1
      let year2 = year
      let month2 = month + monthNum
      if (month2 > 12) {
        year2 += 1
        month2 %= 12
      }
      return new Date (year2 + '-' + ('0' + month2).slice(-2) + '-01')
    },
    minusMonth (date, monthNum) {
      let year = date.getFullYear()
      let month = date.getMonth() + 1
      let year2 = year
      let month2 = month - monthNum
      if (month2 <= 0) {
        year2 -= 1
        month2 += 12
      }
      let dayNum = new Date(year2, month2, 0).getDate()
      return new Date (year2 + '-' + ('0' + month2).slice(-2) + '-' + dayNum + ' 00:00:00')
    },
    getDays (year, month) {
      let dayNum = new Date(year, month, 0).getDate()
      let days = Array.apply(null, new Array(new Date(year, month - 1, 1).getDay())).map(() => '')
      days = [...days, ...Array.apply(null, new Array(dayNum)).map((cur, index) => index + 1)]
      // console.log(days)
      return days
    },
    judgeChosen (year, month, day) {
      let tmpDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
      // console.log(tmpDateStr)
      if (tmpDateStr === this.chooseStartDateStr || tmpDateStr === this.chooseEndDateStr) {
        return true
      }
      return false
    },
    judgeDisable (year, month, day) {
      if (day === '') {
        return true
      }
      let tmpDate = new Date(year + '-' + month + '-' + day)
      if (tmpDate < this.leftDate || tmpDate > this.endDate) {
        return true
      }
      return false
    },
    judgeInRange(year, month, day) {
      if (day === '') {
        return false
      }
      let tmpDate = new Date(year + '-' + month + '-' + day)
      // console.log('.....................')
      // console.log(tmpDate)
      // console.log(new Date(this.chooseStartDateStr))
      // console.log(new Date(this.chooseEndDateStr))
      if (tmpDate >= new Date(this.chooseStartDateStr + ' 00:00:00') && tmpDate <= new Date(this.chooseEndDateStr + ' 00:00:00')) {
        return true
      }
      return false
    },
    judgeStart (year, month, day) {
      let tmpCurDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
      if (this.chooseStartDateStr && this.chooseEndDateStr) {
        // 开始日期和结束日期都选中
        if (tmpCurDateStr === this.chooseStartDateStr) {
          return true
        }
      } else if (this.chooseStartDateStr || this.chooseEndDateStr) {
        // 只选中一个, 要考虑hover的预选日期情况
        let tmpChooseDateStr = this.chooseStartDateStr || this.chooseEndDateStr
        if (!this.preChooseDateStr) {
          return true
        } else {
          if (this.preChooseDateStr >= tmpChooseDateStr) {
            if (tmpCurDateStr === tmpChooseDateStr) {
              return true
            }
          } else {
            if (tmpCurDateStr === this.preChooseDateStr) {
              return true
            }
          }
        }
      }
      return false
    },
    judgeEnd (year, month, day) {
      let tmpCurDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
      if (this.chooseStartDateStr && this.chooseEndDateStr) {
        // 开始日期和结束日期都选中
        if (tmpCurDateStr === this.chooseEndDateStr) {
          return true
        }
      } else if (this.chooseStartDateStr || this.chooseEndDateStr) {
        // 只选中一个, 要考虑hover的预选日期情况
        let tmpChooseDateStr = this.chooseStartDateStr || this.chooseEndDateStr
        if (!this.preChooseDateStr) {
          return true
        } else {
          if (this.preChooseDateStr <= tmpChooseDateStr) {
            if (tmpCurDateStr === tmpChooseDateStr) {
              return true
            }
          } else {
            if (tmpCurDateStr === this.preChooseDateStr) {
              return true
            }
          }
        }
      }
      return false
    },
    judgeToday (year, month, day) {
      let today = new Date()
      if (year === today.getFullYear() && month === today.getMonth() + 1 && day === today.getDate()) {
        return true
      }
      return false
    },
    chooseDay (year, month, day) {
      if (this.judgeDisable(year, month, day)) {
        return
      }
      let tmpDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
      console.log(tmpDateStr)
      if (this.chooseStartDateStr && this.chooseEndDateStr) {
        // 开始和结束日期都已选中
        if (tmpDateStr === this.chooseStartDateStr) {
          this.chooseStartDateStr = this.chooseEndDateStr
          this.chooseEndDateStr = ''
        } else if (tmpDateStr === this.chooseEndDateStr) {
          this.chooseEndDateStr = ''
        } else if (tmpDateStr > this.chooseEndDateStr) {
          this.chooseEndDateStr = tmpDateStr
        } else if (tmpDateStr < this.startDateStr) {
          this.chooseStartDateStr = tmpDateStr
        } else {
          this.chooseEndDateStr = tmpDateStr
        }
      } else if (this.chooseStartDateStr || this.chooseEndDateStr) {
        // 开始和结束日期仅选中一个
        if (tmpDateStr > this.chooseStartDateStr) {
          this.chooseEndDateStr = tmpDateStr
        } else if (tmpDateStr < this.chooseStartDateStr) {
          this.chooseEndDateStr = this.chooseStartDateStr
          this.chooseStartDateStr = tmpDateStr
        } else {
          this.chooseStartDateStr = ''
        }
      } else {
        // 开始和结束日期均未选中
        this.chooseStartDateStr = tmpDateStr
      }
    },
    preChoose (year, month, day) {
      if (!this.judgeDisable(year, month, day)) {
        this.preChooseDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
        // console.log(this.preChooseDateStr)
        this.$forceUpdate()
      }
    },
    preNotChoose (year, month, day) {
      this.preChooseDateStr = ''
    },
    judgeInPreRange (year, month, day) {
      // console.log(year, month, day)
      if (!this.preChooseDateStr) {
        return false
      }
      if (this.judgeDisable(year, month, day)) {
        return false
      }
      if ((this.chooseStartDateStr && this.chooseEndDateStr) || (!this.chooseStartDateStr && !this.chooseEndDateStr)) {
        return false
      } else {
        let tmpCurDateStr = year + '-' + ('0' + month).slice(-2) + '-' + ('0' + day).slice(-2)
        let tmpDateStr = this.chooseStartDateStr || this.chooseEndDateStr
        let left, right
        if (tmpDateStr < this.preChooseDateStr) {
          left = tmpDateStr
          right = this.preChooseDateStr
        } else {
          left = this.preChooseDateStr
          right = tmpDateStr
        }
        console.log(left, right, tmpCurDateStr)
        if (tmpCurDateStr <= right && tmpCurDateStr >= left) {
          // console.log(tmpCurDateStr)
          return true
        } else {
          return false
        }
      }
    },

    goPreMonth () {
      // console.log(new Date(this.curLeftYear, this.curLeftMonth - 1, 1))
      // console.log(this.minusMonth(new Date(this.curLeftYear, this.curLeftMonth - 1, 1),2))
      // console.log(this.leftDate)
      if (this.minusMonth(new Date(this.curLeftYear, this.curLeftMonth - 1, 1), 2) >= this.leftDate ) {
        let curLeftDate = this.minusMonth(new Date(this.curLeftYear, this.curLeftMonth - 1, 1), 2)
        let curRightDate = this.minusMonth(new Date(this.curLeftYear, this.curLeftMonth - 1, 1), 1)
        this.curLeftYear = curLeftDate.getFullYear()
        this.curLeftMonth = curLeftDate.getMonth() + 1
        this.curRightYear = curRightDate.getFullYear()
        this.curRightMonth = curRightDate.getMonth() + 1
        this.curLeftDays = this.getDays(this.curLeftYear, this.curLeftMonth)
        this.curRightDays = this.getDays(this.curRightYear, this.curRightMonth)
      }
    },
    goNextMonth () {
      if (this.addMonth(new Date(this.curRightYear, this.curRightMonth - 1, 1), 2) <= this.rightDate ) {
        let curLeftDate = this.addMonth(new Date(this.curRightYear, this.curRightMonth - 1, 1), 1)
        let curRightDate = this.addMonth(new Date(this.curRightYear, this.curRightMonth - 1, 1), 2)
        this.curLeftYear = curLeftDate.getFullYear()
        this.curLeftMonth = curLeftDate.getMonth() + 1
        this.curRightYear = curRightDate.getFullYear()
        this.curRightMonth = curRightDate.getMonth() + 1
        this.curLeftDays = this.getDays(this.curLeftYear, this.curLeftMonth)
        this.curRightDays = this.getDays(this.curRightYear, this.curRightMonth)
      }
    },
    determine () {
      if (this.chooseStartDateStr && this.chooseEndDateStr) {
        this.$emit('chooseDate', new Date(this.chooseStartDateStr), new Date(this.chooseEndDateStr))
      } else {
        this.$notify.warning({
          title: '提醒',
          message: '开始和结束日期都要选择！'
        })
      }
    },
    cancel () {
      this.$emit('close')
    }
  },
  created() {
    if (this.monthNum % 2) {
      this.monthNum += 1
    }
    let leftDate = new Date(new Date(new Date().getTime()).setHours(0,0,0,0))
    let rightDate = this.addMonth(leftDate, this.monthNum)
    this.leftDate = leftDate
    this.rightDate = rightDate
    console.log(leftDate, rightDate)
    let startDate = new Date(this.startDateStr)
    let endDate = new Date(this.endDateStr)
    console.log(startDate, endDate)
    if (startDate >= endDate || startDate < leftDate || endDate > rightDate) {
      this.chooseStartDateStr = ''
      this.chooseEndDateStr = ''
      // console.log('超范围', this.leftDate)
      this.curLeftYear = this.leftDate.getFullYear()
      this.curLeftMonth = this.leftDate.getMonth() + 1
      this.curRightYear = this.addMonth(this.leftDate, 1).getFullYear()
      this.curRightMonth = this.addMonth(this.leftDate, 1).getMonth() + 1
      this.$notify.warning({
        title: '提醒',
        message: '日期超出范围，请重新选择'
      })
    } else {
      this.chooseStartDateStr = this.startDateStr
      this.chooseEndDateStr = this.endDateStr
      // console.log('没超', this.chooseStartDateStr)
      let chooseStartDate = new Date(this.chooseStartDateStr)
      // console.log(chooseStartDate)
      // let chooseEndDate = new Date(this.chooseEndDateStr)
      this.curLeftYear = chooseStartDate.getFullYear()
      this.curLeftMonth = chooseStartDate.getMonth() + 1
      this.curRightYear = this.addMonth(chooseStartDate, 1).getFullYear()
      this.curRightMonth = this.addMonth(chooseStartDate, 1).getMonth() + 1
      // console.log(this.addMonth(chooseStartDate, 1))
    }
    this.curLeftDays = this.getDays(this.curLeftYear, this.curLeftMonth)
    this.curRightDays = this.getDays(this.curRightYear, this.curRightMonth)
  }
}
</script>

<style lang="less" scoped>
.choose-date-range-components {
  color: #000000;
  display: flex;
  flex-direction: column;
  box-shadow: 0 0 4px 0 rgba(117,117,117,0.5);
  //opacity: 1!important;
  border-radius: 2px;
  .opt-btns {
    padding: 10px;
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid #dddddd;
  }
  .tow-month {
    display: flex;
    .month-bar {
      width: 50%;
      border-right: 1px solid #dddddd;
      padding: 0 10px;
      .header {
        display: flex;
        justify-content: space-between;
        padding: 5px 0;
        .arrow {
          cursor: pointer;
        }
      }
      .weekdays {
        display: flex;
        .weekday {
          width: 40px;
          text-align: center;
        }
      }
      .days {
        display: flex;
        flex-wrap: wrap;
        .day {
          width: 40px;
          text-align: center;
          height: 40px;
          line-height: 30px;
          cursor: pointer;
          padding: 5px;
          .day-inner {
            width: 100%;
            height: 100%;
          }
          .chosen-day {
            //width: 100%;
            //height: 100%;
            border-radius: 20px;
            //border: 1px solid red;
            background-color: #409eff;
            color: #FFFFFF;
          }
        }
        .cannot-choose {
          color: #dddddd;
          cursor: auto;
        }
        .in-range {
          background-color: #f2f6fc;
        }
        .in-pre-range {
          background-color: #f2f6fc;
        }
        .is-start {
          border-radius: 20px 0 0 20px;
        }
        .is-end {
          border-radius: 0 20px 20px 0!important;
        }
        .is-today {
          color: #409eff;
        }
      }
    }
    .month-bar:last-child {
      border-right: none;
    }
  }
}
</style>